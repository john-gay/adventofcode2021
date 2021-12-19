package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type node struct {
	value  int
	x      *node
	y      *node
	parent *node
}

func main() {
	start := time.Now()

	file, err := os.Open("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fishNumbers := []*node{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			fishNumber := &node{}

			for _, char := range scanner.Text() {
				switch char {
				case '[':
					fishNumber.x = &node{parent: fishNumber}
					fishNumber.y = &node{parent: fishNumber}
					fishNumber = fishNumber.x
				case ',':
					fishNumber = fishNumber.parent.y
				case ']':
					fishNumber = fishNumber.parent
				default:
					v, _ := strconv.Atoi(string(char))
					fishNumber.value = v
				}
			}
			fishNumbers = append(fishNumbers, fishNumber)
		}
	}

	// Comment out for part 2 to work
	workingNumber := fishNumbers[0]
	for i := 1; i < len(fishNumbers); i++ {
		workingNumber = workingNumber.add(fishNumbers[i])

		workingNumber.reduce()
	}

	fmt.Println(fmt.Sprintf("Part 1: %d", workingNumber.magnitude()))

	fmt.Println(fmt.Sprintf("Part 2: %d", findLargestMagnitude(fishNumbers)))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (n *node) reduce() {
	reducing := true
	updated := false
	for ok := true; ok; ok = reducing {
		updated = n.reduceAndExplode(0)

		if !updated {
			updated = n.reduceAndSplit()

			if !updated {
				reducing = false
			}
		}
	}
}

func (n *node) reduceAndExplode(depth int) bool {
	if depth > 4 {
		n.parent.explode()
		return true
	}
	if (n.parent == nil && depth > 0) || n.x == nil {
		return false
	}
	updated := n.x.reduceAndExplode(depth + 1)
	if n.y == nil || updated {
		return updated
	}
	return n.y.reduceAndExplode(depth + 1)
}

func (n *node) explode() {
	x := n.findClosestX()
	if x != nil {
		x.value += n.x.value
	}
	y := n.findClosestY()
	if y != nil {
		y.value += n.y.value
	}
	n.x.parent, n.x, n.y.parent, n.y = nil, nil, nil, nil
	n.value = 0
}

func (n *node) findClosestX() *node {
	child := n
	currentNode := n.parent
	for currentNode != nil {
		if currentNode.y == child {
			child = currentNode
			currentNode = currentNode.x
		} else if currentNode.x == child {
			child = currentNode
			currentNode = currentNode.parent
		} else if currentNode.x == nil && currentNode.y == nil {
			return currentNode
		} else {
			currentNode = currentNode.y
		}
	}

	return nil
}

func (n *node) findClosestY() *node {
	child := n
	currentNode := n.parent
	for currentNode != nil {
		if currentNode.x == child {
			child = currentNode
			currentNode = currentNode.y
		} else if currentNode.y == child {
			child = currentNode
			currentNode = currentNode.parent
		} else if currentNode.x == nil && currentNode.y == nil {
			return currentNode
		} else {
			currentNode = currentNode.x
		}
	}

	return nil
}

func (n *node) add(new *node) *node {
	newNode := &node{
		x: n,
		y: new,
	}
	newNode.x.parent = newNode
	newNode.y.parent = newNode

	return newNode
}

func (n *node) reduceAndSplit() bool {
	if n.x == nil && n.y == nil {
		if n.value > 9 {
			n.x = &node{
				x:      nil,
				y:      nil,
				parent: n,
				value:  int(math.Floor(float64(n.value) / 2)),
			}
			n.y = &node{
				x:      nil,
				y:      nil,
				parent: n,
				value:  int(math.Ceil(float64(n.value) / 2)),
			}
			n.value = 0
			return true
		} else {
			return false
		}
	}
	if n.x == nil {
		return false
	}
	updated := n.x.reduceAndSplit()
	if n.y == nil || updated {
		return updated
	}
	return n.y.reduceAndSplit()
}

func (n *node) magnitude() int {
	mag := 0
	if n.x != nil {
		mag += n.x.magnitude() * 3
	}
	if n.y != nil {
		mag += n.y.magnitude() * 2
	}
	if n.x == nil && n.y == nil {
		return n.value
	}

	return mag
}

func (n *node) String() string {
	builder := strings.Builder{}
	if n.x != nil {
		builder.WriteString("[")
		builder.WriteString(n.x.String())
		builder.WriteString(",")
	}
	if n.y != nil {
		builder.WriteString(n.y.String())
		builder.WriteString("]")
	}
	if n.x == nil && n.y == nil {
		builder.WriteString(strconv.Itoa(n.value))
	}

	return builder.String()
}

func (n *node) printParent() {
	if n.parent != nil {
		n.parent.printParent()
	} else {
		fmt.Println(n.String())
	}
}

func findLargestMagnitude(fishNumbers []*node) int {
	max := 0
	for i := 0; i < len(fishNumbers); i++ {
		for j := 0; j < len(fishNumbers); j++ {
			if i != j {
				fishA := fishNumbers[i].copy()
				fishB := fishNumbers[j].copy()

				fishC := fishA.add(fishB)
				fishC.reduce()

				value := fishC.magnitude()
				if value > max {
					max = value
				}
			}
		}
	}
	return max
}

func (n *node) copy() *node {
	c := &node{}
	if n.x != nil {
		c.x = n.x.copy()
		c.x.parent = c
	}
	if n.y != nil {
		c.y = n.y.copy()
		c.y.parent = c
	}
	c.value = n.value

	return c
}

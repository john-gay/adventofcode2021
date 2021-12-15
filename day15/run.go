package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type task struct {
	grid      [][]int
	unvisited map[string]bool
	minDist   map[string]int
	partials  map[string]int
}

func main() {
	start := time.Now()

	file, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	intialGridSize := 100
	gridSize := intialGridSize * 5

	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	unvisited := map[string]bool{}
	minDist := map[string]int{}

	index := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), "")

			for j, part := range parts {
				n, _ := strconv.ParseInt(part, 10, 8)
				grid[index][j] = int(n)
			}
			index++
		}
	}

	grid = makeLargeGrid(intialGridSize, grid)

	for li := range grid {
		for lj := range grid[li] {
			key := fmt.Sprintf("%d-%d", li, lj)
			unvisited[key] = false
			minDist[key] = int(^uint(0) >> 1)
		}
	}

	minDist[fmt.Sprintf("%d-%d", 0, 0)] = 0

	t := task{
		grid:      grid,
		unvisited: unvisited,
		minDist:   minDist,
		partials:  map[string]int{},
	}

	searching := true

	i := 0
	j := 0
	for ok := true; ok; ok = searching {
		t.checkNeighbours(i, j)

		i, j = t.findLowestNotVisited()
		if i < 0 {
			searching = false
		}
	}

	fmt.Println(fmt.Sprintf("Part 1: %d", t.minDist["99-99"]))

	fmt.Println(fmt.Sprintf("Part 2: %d", t.minDist["499-499"]))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func makeLargeGrid(gridSize int, grid [][]int) [][]int {
	for n := 1; n < 5; n++ {
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				if j+(gridSize*n) == 50 {
					fmt.Println("oh no")
				}
				grid[i][j+(gridSize*n)] = newGridValue(grid[i][j], n)
			}
		}
	}
	for n := 1; n < 5; n++ {
		for i := 0; i < gridSize; i++ {
			for j := range grid[0] {
				grid[i+(gridSize*n)][j] = newGridValue(grid[i][j], n)
			}
		}
	}
	return grid
}

func newGridValue(i, n int) int {
	v := i + n

	if v >= 10 {
		v -= 9
	}

	return v
}

func (t *task) checkNeighbours(i, j int) {
	delete(t.partials, fmt.Sprintf("%d-%d", i, j))
	neighbours := [][]int{{i + 1, j}, {i, j + 1}, {i - 1, j}, {i, j - 1}}

	for _, neighbour := range neighbours {
		key := fmt.Sprintf("%d-%d", neighbour[0], neighbour[1])
		if visited, ok := t.unvisited[key]; ok {
			if !visited {
				newValue := t.minDist[fmt.Sprintf("%d-%d", i, j)] + t.grid[neighbour[0]][neighbour[1]]
				if newValue < t.minDist[key] {
					t.minDist[key] = newValue
					t.partials[key] = newValue
				}
			}
		}
	}

	t.unvisited[fmt.Sprintf("%d-%d", i, j)] = true
}

func (t *task) findLowestNotVisited() (int, int) {
	minKey := ""
	minValue := int(^uint(0) >> 1)
	for key, value := range t.partials {
		if !t.unvisited[key] && minValue > value {
			minValue = value
			minKey = key
		}
	}

	if minKey != "" {
		parts := strings.Split(minKey, "-")
		newI, _ := strconv.ParseInt(parts[0], 10, 64)
		newJ, _ := strconv.ParseInt(parts[1], 10, 64)

		return int(newI), int(newJ)
	}
	return -1, -1

}

func printGrid(grid [][]int) {
	fmt.Println("")
	for _, row := range grid {
		fmt.Println(row)
	}
}

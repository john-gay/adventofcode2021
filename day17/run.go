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

type bucket struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func main() {
	start := time.Now()

	file, err := os.Open("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var b bucket

	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), ", y=")
			yParts := strings.Split(parts[1], "..")
			parts = strings.Split(parts[0], ": x=")
			xParts := strings.Split(parts[1], "..")
			xMin, _ := strconv.ParseInt(xParts[0], 10, 32)
			xMax, _ := strconv.ParseInt(xParts[1], 10, 32)
			yMin, _ := strconv.ParseInt(yParts[0], 10, 32)
			yMax, _ := strconv.ParseInt(yParts[1], 10, 32)
			b = bucket{
				xMin: int(xMin),
				xMax: int(xMax),
				yMin: int(yMin),
				yMax: int(yMax),
			}
		}
	}

	// Max y velocity is yMin - 1
	// Formula to sum series x-1 is ((x+1) ^ 2 - (x+1)) / 2
	highest := (b.yMin*b.yMin - b.yMin) / 2

	fmt.Println(fmt.Sprintf("Part 1: %d", highest))
	fmt.Println(fmt.Sprintf("Part 2: %d", b.numOfInitialVelocities()))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (b *bucket) numOfInitialVelocities() int {
	total := 0

	for x := 1; x <= b.xMax; x++ {
		for y := b.yMin; y <= b.yMin*-1; y++ {
			if b.landInBucket(x, y) {
				total++
			}
		}
	}

	return total
}

func (b *bucket) landInBucket(x int, y int) bool {
	// Return false if x will never reach minimum distance
	if ((x+1)*(x+1)-(x+1))/2 < b.xMin {
		return false
	}

	// Any value of x between half of max x value and less than min x
	// will step too far on the second step
	if x <= b.xMin && x >= b.xMax+1/2 {
		return false
	}

	plotting := true
	inBucket := false
	yv, yd, xv, xd := y, y, x, x
	for ok := true; ok; ok = plotting {
		if xd <= b.xMax && xd >= b.xMin && yd <= b.yMax && yd >= b.yMin {
			inBucket = true
			plotting = false
		}
		if xd >= b.xMax+1 || yd <= b.yMin-1 {
			plotting = false
		}
		yd, yv = yDistance(yv, yd)
		xd, xv = xDistance(xv, xd)
	}
	return inBucket
}

func xDistance(v, d int) (int, int) {
	if v < 1 {
		return d, v
	}
	return d + v - 1, v - 1
}

func yDistance(v, d int) (int, int) {
	return d + v - 1, v - 1
}

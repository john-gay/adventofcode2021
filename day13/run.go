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

func main() {
	start := time.Now()

	file, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := [][]int{}
	folds := [][]string{}

	maxX := 0
	maxY := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), ",")
			point := []int{}
			for _, part := range parts {
				d, err := strconv.ParseInt(part, 10, 32)
				if err != nil {
					foldParts := strings.Split(part, " ")
					coordParts := strings.Split(foldParts[2], "=")
					folds = append(folds, coordParts)
				} else {
					point = append(point, int(d))
				}
			}
			if len(point) > 0 {
				points = append(points, point)
				if maxX < point[0] {
					maxX = point[0]
				}
				if maxY < point[1] {
					maxY = point[1]
				}
			}
		}
	}

	grid := make([][]string, maxY+1)
	for i := range grid {
		grid[i] = make([]string, maxX+1)
	}

	plotGrid(grid, points)

	printGrid(grid)

	part1Count := 0

	for i, fold := range folds {
		grid = foldGrid(grid, fold)

		if i == 0 {
			part1Count = countGrid(grid)
		}

		printGrid(grid)
	}

	fmt.Println(fmt.Sprintf("Part 1: %d", part1Count))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func countGrid(grid [][]string) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				count++
			}
		}
	}
	return count
}

func foldGrid(grid [][]string, fold []string) [][]string {
	if fold[0] == "y" {
		f, _ := strconv.ParseInt(fold[1], 10, 32)
		foldLine := int(f)
		for i := foldLine + 1; i < len(grid); i++ {
			for j := range grid[i] {
				if grid[i][j] == "#" {
					newLine := foldLine - (i - foldLine)
					grid[newLine][j] = "#"
				}
			}
		}
		grid = grid[:foldLine]
	}
	if fold[0] == "x" {
		f, _ := strconv.ParseInt(fold[1], 10, 32)
		foldLine := int(f)
		for j := foldLine + 1; j < len(grid[0]); j++ {
			for i := range grid {
				if grid[i][j] == "#" {
					newLine := foldLine - (j - foldLine)
					grid[i][newLine] = "#"
				}
			}
		}
		for i := range grid {
			grid[i] = grid[i][:foldLine]
		}
	}
	return grid
}

func plotGrid(grid [][]string, points [][]int) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = "_"
		}
	}
	for _, point := range points {
		grid[point[1]][point[0]] = "#"
	}
}

func printGrid(grid [][]string) {
	fmt.Println("")
	for _, row := range grid {
		fmt.Println(row)
	}
}

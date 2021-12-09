package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int
	var row []int

	nineRow := []int{}
	partSize := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), "")
			if partSize == 0 {
				partSize = len(parts) + 2
			}
			row = append(row, 9)
			for _, part := range parts {
				digit, _ := strconv.ParseInt(part, 10, 8)
				row = append(row, int(digit))
			}
			row = append(row, 9)
			grid = append(grid, row)
			row = []int{}
		}
	}

	for i := 0; i < partSize; i++ {
		nineRow = append(nineRow, 9)
	}
	grid = append([][]int{nineRow}, grid...)
	grid = append(grid, nineRow)

	basinSizes := []int{}

	for i := range grid {
		for j, cell := range grid[i] {
			if cell == 9 {
				continue
			}

			if grid[i-1][j] > cell && grid[i+1][j] > cell && grid[i][j-1] > cell && grid[i][j+1] > cell {
				basinSizes = append(basinSizes, calculateBasinSize(i, j, grid))
			}
		}
	}

	sort.Ints(basinSizes)

	largestThree := basinSizes[len(basinSizes)-3:]

	total := 1
	for _, size := range largestThree {
		total *= size
	}

	fmt.Println(total)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func calculateBasinSize(i, j int, grid [][]int) int {
	cellMap := make(map[string]bool)
	cellMap[fmt.Sprintf("%d+%d", i, j)] = true

	exploreSurrounding(i, j, grid, cellMap)

	return len(cellMap)
}

func exploreSurrounding(i int, j int, grid [][]int, cellMap map[string]bool) {
	if grid[i+1][j] != 9 {
		exploreIfNotExplored(i+1, j, grid, cellMap)
	}
	if grid[i-1][j] != 9 {
		exploreIfNotExplored(i-1, j, grid, cellMap)
	}
	if grid[i][j+1] != 9 {
		exploreIfNotExplored(i, j+1, grid, cellMap)
	}
	if grid[i][j-1] != 9 {
		exploreIfNotExplored(i, j-1, grid, cellMap)
	}
}

func exploreIfNotExplored(i int, j int, grid [][]int, cellMap map[string]bool) {
	_, present := cellMap[fmt.Sprintf("%d+%d", i, j)]
	if !present {
		cellMap[fmt.Sprintf("%d+%d", i, j)] = true
		exploreSurrounding(i, j, grid, cellMap)
	}
}

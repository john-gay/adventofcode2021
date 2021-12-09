package main

import (
	"bufio"
	"log"
	"os"
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

	risk := 0

	for i := range grid {
		for j, cell := range grid[i] {
			if cell == 9 {
				continue
			}

			if grid[i-1][j] > cell && grid[i+1][j] > cell && grid[i][j-1] > cell && grid[i][j+1] > cell {
				risk += cell + 1
			}
		}
	}

	log.Println(risk)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

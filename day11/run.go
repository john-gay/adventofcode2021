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

	file, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var octs [][]int

	for scanner.Scan() {
		row := []int{}
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), "")
			for _, part := range parts {
				d, _ := strconv.ParseInt(part, 10, 8)
				row = append(row, int(d))
			}
			octs = append(octs, row)
		}
	}

	flashes := 0

	turn := 0
	syncing := true

	for ok := true; ok; ok = syncing {
		turn++

		increaseByOne(octs)

		flashOcts(octs)

		turnFlashes := resetOcts(octs)
		flashes += turnFlashes

		if turnFlashes == 100 {
			syncing = false
		}

		if turn == 100 {
			fmt.Println(fmt.Sprintf("Part 1: %d", flashes))
		}
	}

	fmt.Println(fmt.Sprintf("Part 2: %d", turn))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func flashOcts(octs [][]int) {
	flashed := map[string]bool{}

	for i, _ := range octs {
		for j, _ := range octs[i] {
			if octs[i][j] > 9 {
				flashOct(i, j, octs, flashed)
			}
		}
	}
}

func flashOct(i int, j int, octs [][]int, flashed map[string]bool) {
	if validOct(i, j, octs) && !flashed[fmt.Sprintf("%d-%d", i, j)] {
		flashed[fmt.Sprintf("%d-%d", i, j)] = true

		increaseOct(i+1, j, octs, flashed)
		increaseOct(i+1, j+1, octs, flashed)
		increaseOct(i, j+1, octs, flashed)
		increaseOct(i-1, j, octs, flashed)
		increaseOct(i-1, j-1, octs, flashed)
		increaseOct(i, j-1, octs, flashed)
		increaseOct(i+1, j-1, octs, flashed)
		increaseOct(i-1, j+1, octs, flashed)
	}
}

func increaseOct(i int, j int, octs [][]int, flashed map[string]bool) {
	if validOct(i, j, octs) {
		octs[i][j]++
		if octs[i][j] > 9 {
			flashOct(i, j, octs, flashed)
		}
	}
}

func validOct(i int, j int, octs [][]int) bool {
	return i >= 0 && j >= 0 && i < len(octs) && j < len(octs[0])
}

func resetOcts(octs [][]int) int {
	flashes := 0
	for i, _ := range octs {
		for j, _ := range octs[i] {
			if octs[i][j] > 9 {
				octs[i][j] = 0
				flashes++
			}
		}
	}
	return flashes
}

func increaseByOne(octs [][]int) {
	for i, _ := range octs {
		for j, _ := range octs[i] {
			octs[i][j]++
		}
	}
}

func showOcts(octs [][]int) {
	fmt.Println("")
	for _, row := range octs {
		fmt.Println(row)
	}
}

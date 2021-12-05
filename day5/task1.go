package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	start point
	end   point
}

type point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []line

	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), " -> ")
			startParts := strings.Split(parts[0], ",")
			startX, _ := strconv.ParseInt(startParts[0], 10, 32)
			startY, _ := strconv.ParseInt(startParts[1], 10, 32)
			endParts := strings.Split(parts[1], ",")
			endX, _ := strconv.ParseInt(endParts[0], 10, 32)
			endY, _ := strconv.ParseInt(endParts[1], 10, 32)
			lines = append(lines, line{
				start: point{
					x: int(startX),
					y: int(startY),
				},
				end: point{
					x: int(endX),
					y: int(endY),
				},
			})
		}
	}

	fmt.Println(fmt.Sprintf("Lines: %v", lines))

	processLines(lines)
}

func processLines(lines []line) {
	board := drawBoard(1000, 1000)

	board = drawLines(board, lines)

	fmt.Println(countOverlaps(board))
}

func countOverlaps(board [][]int) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] >= 2 {
				count++
			}
		}
	}
	return count
}

func drawBoard(maxX, mayY int) [][]int {
	board := make([][]int, maxX)
	for i := range board {
		board[i] = make([]int, mayY)
	}
	return board
}

func drawLines(board [][]int, lines []line) [][]int {
	for _, line := range lines {
		if (line.start.x == line.end.x) || (line.start.y == line.end.y) {
			board = drawLine(board, line)
		}
	}

	return board
}

func drawLine(board [][]int, line line) [][]int {
	if line.start.x == line.end.x {
		x := line.start.x
		var minY int
		var maxY int
		if line.start.y < line.end.y {
			minY = line.start.y
			maxY = line.end.y
		} else {
			minY = line.end.y
			maxY = line.start.y
		}
		for y := minY; y <= maxY; y++ {
			board[x][y]++
		}
	} else {
		y := line.start.y
		var minX int
		var maxX int
		if line.start.x < line.end.x {
			minX = line.start.x
			maxX = line.end.x
		} else {
			minX = line.end.x
			maxX = line.start.x
		}
		for x := minX; x <= maxX; x++ {
			board[x][y]++
		}
	}
	return board
}

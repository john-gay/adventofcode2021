package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var turns []string
	var board [][]string
	var boards [][][]string

	rowCount := 0
	boardRow := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			var cell strings.Builder
			var turn strings.Builder
			var row []string
			for _, char := range scanner.Text() {
				if rowCount == 0 {
					if string(char) == "," {
						turns = append(turns, turn.String())
						turn = strings.Builder{}
					} else {
						turn.WriteString(string(char))
					}
				} else {
					if string(char) != " " {
						cell.WriteString(string(char))
					} else {
						if cell.String() != "" {
							row = append(row, cell.String())
							cell = strings.Builder{}
						}
					}
				}
			}
			if turn.String() != "" {
				turns = append(turns, turn.String())
				turn = strings.Builder{}
			}

			if cell.String() != "" {
				row = append(row, cell.String())
			}
			if len(row) > 0 {
				board = append(board, row)
				boardRow++
			}

			if boardRow == 5 {
				boards = append(boards, board)
				board = [][]string{}
				boardRow = 0
			}
		}

		rowCount++
	}

	fmt.Println(fmt.Sprintf("Turns: %v", turns))
	fmt.Println(fmt.Sprintf("Boards: %v", boards))

	playBingo(turns, boards)
}

func playBingo(turns []string, boards [][][]string) {
	scoreBoards := make([][][]uint8, len(boards))
	for i := range scoreBoards {
		scoreBoards[i] = make([][]uint8, len(boards[0]))
		for j := range scoreBoards[i] {
			scoreBoards[i][j] = make([]uint8, len(boards[0][0]))
		}
	}

	var winners []int

	for _, turn := range turns {
		for i, board := range boards {
			if !hasBoardWon(i, winners) {
				scoreBoards[i] = makeTurn(turn, board, scoreBoards[i])
				winner := checkBoard(scoreBoards[i])
				if winner == true {
					winners = append(winners, i)
				}
			}
			if len(winners) == len(boards) {
				fmt.Println(fmt.Sprintf("Last winner: %v", scoreBoards[i]))

				fmt.Println(calculateScore(turn, board, scoreBoards[i]))
				return
			}
		}
	}

	fmt.Println(fmt.Sprintf("Score finished: %v", scoreBoards))
}

func hasBoardWon(board int, winners []int) bool {
	for _, winner := range winners {
		if winner == board {
			return true
		}
	}
	return false
}

func calculateScore(turn string, board [][]string, scoreBoard [][]uint8) int {
	sum := 0
	for i := range board {
		for j := range board[i] {
			if scoreBoard[i][j] == 0 {
				cell, _ := strconv.Atoi(board[i][j])
				sum += cell
			}
		}
	}
	turnInt, _ := strconv.Atoi(turn)
	return sum * turnInt
}

func checkBoard(scoreBoard [][]uint8) bool {
	rowSize := len(scoreBoard)
	colSize := len(scoreBoard[0])

	for i := 0; i < rowSize; i++ {
		winRow := true
		for j := 0; j < colSize; j++ {
			if scoreBoard[i][j] == 0 {
				winRow = false
			}
		}
		if winRow {
			return true
		}
	}

	for j := 0; j < colSize; j++ {
		winCol := true
		for i := 0; i < rowSize; i++ {
			if scoreBoard[i][j] == 0 {
				winCol = false
			}
		}
		if winCol {
			return true
		}
	}

	return false
}

func makeTurn(turn string, board [][]string, scoreBoard [][]uint8) [][]uint8 {
	for i := range board {
		for j := range board[i] {
			if turn == board[i][j] {
				scoreBoard[i][j] = 1
			}
		}
	}
	return scoreBoard
}

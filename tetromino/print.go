package tetromino

import (
	"math"
	"strings"
)

// pos represents a coordinate on the board.
type pos struct {
	X int
	Y int
}

// AssembleTetrominoes arranges the tetrominoes on a square board.
func AssembleTetrominoes(tetrominoes map[int]string) [][]rune {
	size := CalculateMinimumSquareSize(len(tetrominoes))
	board := initializeBoard(size)

	for {
		allPlaced := true

		for i := 1; i <= len(tetrominoes); i++ {
			placed := false

			for x := 0; x < size && !placed; x++ {
				for y := 0; y < size && !placed; y++ {
					if CanPutTetromino(tetrominoes[i], board, x, y) {
						board = PutTetromino(tetrominoes[i], board, x, y)
						printTetromino(tetrominoes[i]) // Print the tetromino being placed
						placed = true
					}
				}
			}

			if !placed {
				allPlaced = false
			}
		}

		if !allPlaced {
			// Expand the board and restart the process
			size++
			board = initializeBoard(size)
		} else {
			break // Exit if all tetrominoes are placed
		}
	}

	return board
}

// printTetromino prints the tetromino to the console.
func printTetromino(tetromino string) {
	lines := strings.Split(tetromino, "\n")
	for _, line := range lines {
		if line != "" {
			println(line)
		}
	}
	println() // Print a new line for separation
}

// initializeBoard creates a square board filled with dots.
func initializeBoard(size int) [][]rune {
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.' // Fill with dots
		}
	}
	return board
}

// PutTetromino places a tetromino on the board at the specified position, maintaining its shape.
func PutTetromino(tetromino string, board [][]rune, x int, y int) [][]rune {
	lines := strings.Split(tetromino, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char >= 'A' && char <= 'Z' {
				board[x+i][y+j] = char // Assign the letter directly
			}
		}
	}
	return board
}

// CanPutTetromino checks if a tetromino can be placed at the given position on the board.
func CanPutTetromino(tetromino string, board [][]rune, x int, y int) bool {
	lines := strings.Split(tetromino, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char >= 'A' && char <= 'Z' {
				if x+i >= len(board) || y+j >= len(board) || board[x+i][y+j] != '.' {
					return false
				}
			}
		}
	}
	return true
}

// CalculateMinimumSquareSize calculates the minimum square size needed for the given number of tetrominoes.
func CalculateMinimumSquareSize(numTetrominoes int) int {
	return int(math.Ceil(math.Sqrt(float64(numTetrominoes * 4))))
}

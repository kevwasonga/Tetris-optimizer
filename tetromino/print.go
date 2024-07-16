package tetromino

import (
	"math"
	"strings"
)

// AssembleTetrominoes arranges the tetrominoes on a square board.
func AssembleTetrominoes(tetrominoes map[rune]string) [][]rune {
	size := CalculateMinimumSquareSize(len(tetrominoes))

	for {
		board := initializeBoard(size)
		if placeTetrominoes(tetrominoes, board, 'A', size) {
			return board
		}
		size++
	}
}

// placeTetrominoes attempts to place all tetrominoes on the board recursively.
func placeTetrominoes(tetrominoes map[rune]string, board [][]rune, current rune, size int) bool {
	if current > 'A'+rune(len(tetrominoes)-1) {
		return true // All tetrominoes placed
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if CanPutTetromino(tetrominoes[current], board, x, y) {
				// Place the tetromino
				PutTetromino(tetrominoes[current], board, x, y, current)

				// Recursive call to place the next tetromino
				if placeTetrominoes(tetrominoes, board, current+1, size) {
					return true
				}

				// Backtrack: Remove the tetromino
				RemoveTetromino(tetrominoes[current], board, x, y)
			}
			// If can't place, skip to next cell immediately
		}
	}

	return false // No valid placement found
}

// RemoveTetromino removes the tetromino from the board at the specified position.
func RemoveTetromino(tetromino string, board [][]rune, x int, y int) {
	lines := strings.Split(tetromino, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				board[x+i][y+j] = '.' // Clear the position
			}
		}
	}
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
func PutTetromino(tetromino string, board [][]rune, x int, y int, key rune) {
	lines := strings.Split(tetromino, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				board[x+i][y+j] = key // Assign the letter directly based on key
			}
		}
	}
}

// CanPutTetromino checks if a tetromino can be placed at the given position on the board.
func CanPutTetromino(tetromino string, board [][]rune, x int, y int) bool {
	lines := strings.Split(tetromino, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
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

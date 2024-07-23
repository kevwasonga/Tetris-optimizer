package tetromino

import (
	"math"
)

// AssembleTetrominoes arranges the tetrominoes on a square board.
func AssembleTetrominoes(tetrominoes map[rune][][]rune) [][]rune {
	size := CalculateMinimumSquareSize(len(tetrominoes))
	board := initializeBoard(size)

	for {
		if placeTetrominoes(tetrominoes, board, 'A', size) {
			return board
		}
		size++
		board = initializeBoard(size) // Reinitialize the board with new size
	}
}

// placeTetrominoes attempts to place all tetrominoes on the board recursively.
func placeTetrominoes(tetrominoes map[rune][][]rune, board [][]rune, current rune, size int) bool {
	if current > 'A'+rune(len(tetrominoes)-1) {
		return true // All tetrominoes placed
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if CanPutTetromino(tetrominoes[current], board, x, y) {
				PutTetromino(tetrominoes[current], board, x, y, current)
				if placeTetrominoes(tetrominoes, board, current+1, size) {
					return true
				}
				RemoveTetromino(tetrominoes[current], board, x, y)
			}
		}
	}

	return false // No valid placement found
}

// RemoveTetromino removes the tetromino from the board at the specified position.
func RemoveTetromino(tetromino [][]rune, board [][]rune, x int, y int) {
	for i, row := range tetromino {
		for j, char := range row {
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
func PutTetromino(tetromino [][]rune, board [][]rune, x int, y int, key rune) {
	for i, row := range tetromino {
		for j, char := range row {
			if char == '#' {
				board[x+i][y+j] = key // Assign the letter directly based on key
			}
		}
	}
}

// CanPutTetromino checks if a tetromino can be placed at the given position on the board.
func CanPutTetromino(tetromino [][]rune, board [][]rune, x int, y int) bool {
	for i, row := range tetromino {
		for j, char := range row {
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

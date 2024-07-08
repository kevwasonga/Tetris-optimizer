package tetromino

import "strings"

func AssembleTetrominoes(tetrominoes map[int]string) [][]rune {
	size := calculateMinimumSquareSize(len(tetrominoes))
	square := make([][]rune, size)
	for i := range square {
		square[i] = make([]rune, size)
		for j := range square[i] {
			square[i][j] = ' ' // Fill with spaces
		}
	}

	// Simplified placement without rotation
	x, y := 0, 0
	for _, tetromino := range tetrominoes {
		lines := strings.Split(tetromino, "\n")
		for i, line := range lines {
			for j, char := range line {
				square[y+i][x+j] = char
			}
		}
		x += 4
		if x >= size {
			x = 0
			y += 4
		}
	}
	return square
}

func calculateMinimumSquareSize(numTetrominoes int) int {
	// Each tetromino is 4x4, so we need to fit numTetrominoes * 16 cells
	totalCells := numTetrominoes * 16
	size := 1
	for size*size < totalCells {
		size++
	}
	return size
}

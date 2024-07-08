package main

import (
	"fmt"
	"strings"

	"tetris/tetromino"
)

func main() {
	bannerMapmap := tetromino.Loadbanner
	tetrominoes, err := bannerMapmap("sample.txt")
	if err != nil {
		fmt.Println("Error loading tetrominoes:", err)
		return
	}

	valid := tetromino.IsValidTetromino

	validTetrominoes := make(map[int]string)
	for key, tetrominoe := range tetrominoes {
		if valid(tetrominoe) {
			validTetrominoes[key] = tetromino.ReplaceChars(tetrominoe, key)
		} else {
			fmt.Printf("Tetromino %d is invalid:\n%s\n", key, tetrominoe)
		}
	}

	if len(validTetrominoes) == 0 {
		fmt.Println("No valid tetrominoes found.")
		return
	}

	square := assembleTetrominoes(validTetrominoes)
	fmt.Println("Assembled Tetrominoes into Square:")
	printSquare(square)
}

func assembleTetrominoes(tetrominoes map[int]string) [][]rune {
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

func printSquare(square [][]rune) {
	for _, row := range square {
		fmt.Println(string(row))
	}
}

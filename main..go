package main

import (
	"fmt"

	"tetromino/tetromino"
)

func main() {
	fileName := "sample.txt" // Replace with your actual sample file path
	bannerMap, err := tetromino.Loadbanner(fileName)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// Temporary map to store valid tetrominos
	validTetrominoes := make(map[string]struct{})

	// Validate all tetrominos first
	for _, tetrominoStr := range bannerMap {
		if tetromino.IsValidTetromino(tetrominoStr) {
			validTetrominoes[tetrominoStr] = struct{}{}
		} else {
			fmt.Println("ERROR")
			return
		}
	}

	// Map valid tetrominos to their corresponding letters
	finalTetrominoes := make(map[rune]string)
	for key, tetrominoStr := range bannerMap {
		if _, exists := validTetrominoes[tetrominoStr]; exists {
			finalTetrominoes[rune('A'+key-1)] = tetrominoStr
		}
	}

	// Assemble the valid tetrominoes into square
	square := tetromino.AssembleTetrominoes(finalTetrominoes)

	// Print assembled square
	fmt.Println("Assembled Tetrominoes into Square:")
	printAssembled(square)
}

func printAssembled(square [][]rune) {
	for _, row := range square {
		for _, cell := range row {
			if cell == '.' {
				fmt.Print(".") // Print '.' for empty cells
			} else {
				fmt.Print(string(cell)) // Print the tetromino character
			}
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"log"

	"tetromino/tetromino"
)

func main() {
	fileName := "sample.txt" // Replace with your actual sample file path
	bannerMap, err := tetromino.Loadbanner(fileName)
	if err != nil {
		log.Fatalf("Error loading banner: %v", err)
		return
	}

	validTetrominoes := make(map[int]string)
	for key, tetrominoStr := range bannerMap {
		// Validate the tetromino
		if tetromino.IsValidTetromino(tetrominoStr) {
			// Replace the # with their respective Latin equivalents
			validTetrominoes[key] = tetromino.ReplaceChars(tetrominoStr, key)
		} else {
			fmt.Printf("Tetromino %d is invalid.\n", key)
		}
	}

	// Assemble the valid tetrominoes into square
	square := tetromino.AssembleTetrominoes(validTetrominoes)

	// Print assembled square
	fmt.Println("Assembled Tetrominoes into Square:")
	printAssembled(square)
}

func printAssembled(square [][]rune) {
	for _, row := range square {
		for _, cell := range row {
			if cell == ' ' {
				fmt.Print(" ") // Print '.' for empty cells
			} else {
				fmt.Print(string(cell)) // Print the tetromino character
			}
		}
		fmt.Println()
	}
}

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
	}

	validTetrominoes := make(map[int]string)
	for key, tetrominoStr := range bannerMap {
		// Validate the tetromino
		if tetromino.IsValidTetromino(tetrominoStr) {
			// repace the # with  their respective latin equivalents
			validTetrominoes[key] = tetromino.ReplaceChars(tetrominoStr, key)
			//	validTetrominoes[key] = tetrominoStr
		} else {
			fmt.Printf("Tetromino %d is invalid.\n", key)
		}
	}

	// Assemble the valid tetrominoes into square
	square := tetromino.AssembleTetrominoes(validTetrominoes)

	// Print assembled square
	fmt.Println("Assembled Tetrominoes into Square:")
	printAssembled(square, validTetrominoes)
}

func printAssembled(square [][]rune, validTetrominoes map[int]string) {
	// Calculate the minimum square size based on the number of valid tetrominoes
	size := calculateMinimumSquareSize(len(validTetrominoes))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			char := '.'
			if i < len(square) && j < len(square[i]) {
				char = square[i][j]
			}
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func calculateMinimumSquareSize(numTetrominoes int) int {
    // Calculate the total number of cells required to accommodate all tetrominoes
    requiredCells := numTetrominoes * 16 // Each tetromino is 4x4 = 16 cells

    // Calculate the initial smallest square size that can fit all tetrominoes
    size := 1// int(math.Ceil(math.Sqrt(float64(requiredCells))))

    // Increment the size until it can accommodate all tetrominoes
    for {
        totalCells := size * size
        if totalCells >= requiredCells {
            break
        }
        size++
    }

    return size


}

package tetromino

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Loadbanner loads the banner file containing tetrominoes and processes each tetromino.
func Loadbanner(fileName string) (map[int]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR: Failed to open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int]string)
	key := 1
	var tetrominoLines []string
	consecutiveNonEmptyLines := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Check for consecutive non-empty lines
		if line != "" {

			if len(line) != 4 {
				fmt.Println("ERROR")
				os.Exit(0)

			}
			consecutiveNonEmptyLines++
			if consecutiveNonEmptyLines > 4 {
				fmt.Println("ERROR")
				os.Exit(0)
			}

			tetrominoLines = append(tetrominoLines, line)
		} else {
			consecutiveNonEmptyLines = 0
		}

		// If we have 4 lines, process the tetromino
		if len(tetrominoLines) == 4 {
			processedTetromino := processTetromino(tetrominoLines)
			bannerMap[key] = processedTetromino

			key++
			tetrominoLines = nil
		}
	}

	// Handle the case where the file does not end with an empty line
	if len(tetrominoLines) > 0 {
		if len(tetrominoLines) != 4 {
			fmt.Println("ERROR!")
			os.Exit(0)
		}
		processedTetromino := processTetromino(tetrominoLines)
		bannerMap[key] = processedTetromino
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(0)
	}

	return bannerMap, nil
}

// processTetromino processes a single tetromino (4x4 grid) string.
func processTetromino(tetromino []string) string {
	grid := make([][]rune, 4)
	for i, line := range tetromino {
		grid[i] = []rune(line)
	}

	// Remove entirely '.' rows
	var trimmedRows [][]rune
	for _, row := range grid {
		if !isRowDots(row) {
			trimmedRows = append(trimmedRows, row)
		}
	}

	// Remove entirely '.' columns
	var trimmedGrid [][]rune
	for col := 0; col < 4; col++ {
		if !isColumnDots(trimmedRows, col) {
			var column []rune
			for _, row := range trimmedRows {
				column = append(column, row[col])
			}
			trimmedGrid = append(trimmedGrid, column)
		}
	}

	// Convert trimmed grid back to string
	var processedTetromino strings.Builder
	for i := 0; i < len(trimmedRows); i++ {
		for j := 0; j < len(trimmedGrid); j++ {
			if j < len(trimmedRows[i]) {
				processedTetromino.WriteRune(trimmedGrid[j][i])
			} else {
				processedTetromino.WriteRune('.')
			}
		}
		processedTetromino.WriteRune('\n')
	}

	return processedTetromino.String()
}

// isRowDots checks if a row is entirely made of '.'
func isRowDots(row []rune) bool {
	for _, char := range row {
		if char != '.' {
			return false
		}
	}
	return true
}

// isColumnDots checks if a column in the grid is entirely made of '.'
func isColumnDots(grid [][]rune, col int) bool {
	for _, row := range grid {
		if row[col] != '.' {
			return false
		}
	}
	return true
}

package tetromino

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Loadbanner loads the banner file containing tetrominoes and processes each tetromino.
// It returns a map with keys representing the tetromino index and values as the processed tetromino string.
func Loadbanner(fileName string) (map[int]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening file %s: %v", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int]string)
	key := 1
	var tetrominoLines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			tetrominoLines = append(tetrominoLines, line)
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
			return nil, fmt.Errorf("ERROR: Incomplete tetromino, expected 4 lines but got %d", len(tetrominoLines))
		}
		processedTetromino := processTetromino(tetrominoLines)
		bannerMap[key] = processedTetromino
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file %s: %v", fileName, err)
	}

	return bannerMap, nil
}

// processTetromino processes a single tetromino (4x4 grid) string.
// It removes rows and columns made entirely of '.' and returns the processed tetromino string.
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
	for _, row := range trimmedGrid {
		for _, char := range row {
			processedTetromino.WriteRune(char)
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

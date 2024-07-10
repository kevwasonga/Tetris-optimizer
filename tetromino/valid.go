package tetromino

import (
	"fmt"
	"strings"
)

// IsValidTetromino validates tetrominoes based on hash count and connections
func IsValidTetromino(tetromino string) bool {
	// Trim leading and trailing whitespace
	tetromino = strings.TrimSpace(tetromino)

	// Split tetromino into lines
	lines := strings.Split(tetromino, "\n")

	// Validate hash count and connections
	countHashes := 0
	totalConnections := 0
	directions := []struct{ x, y int }{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // right, down, left, up
	}

	// Check each line for validity
	for i, line := range lines {
		// Check characters in line
		for j, char := range line {
			if char != '.' && char != '#' {
				fmt.Printf("Invalid character '%c' in line %d, position %d.\n", char, i+1, j+1)
				return false
			}
			if char == '#' {
				countHashes++
				connections := 0
				// Check connections in all directions
				for _, dir := range directions {
					ni, nj := i+dir.x, j+dir.y
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(line) && lines[ni][nj] == '#' {
						connections++
					}
				}
				if connections == 0 {
					fmt.Printf("No connections found for '#' character at line %d, position %d.\n", i+1, j+1)
					return false
				}
				totalConnections += connections // Accumulate total connections
			}
		}
	}

	// Check if there are exactly 4 '#' characters and total connections are at least 6
	if countHashes != 4 || totalConnections < 6 {
		fmt.Println("Invalid number of '#' characters or insufficient connections.")
		return false
	}

	return true
}

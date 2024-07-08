package tetromino

import "strings"

// IsValidTetromino validates tetrominoes
func IsValidTetromino(tetromino string) bool {
	lines := strings.Split(tetromino, "\n")

	countHashes := 0
	totalConnections := 0
	directions := []struct{ x, y int }{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // right, down, left, up
	}

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				countHashes++
				connections := 0
				for _, dir := range directions {
					ni, nj := i+dir.x, j+dir.y
					if ni >= 0 && ni < 4 && nj >= 0 && nj < 4 && lines[ni][nj] == '#' {
						connections++
					}
				}
				if connections == 0 {
					return false
				}
				totalConnections += connections
			}
		}
	}

	// Check if there are exactly 4 '#' characters
	if countHashes != 4 {
		return false
	}

	// Check if the total connections are at least 6
	if totalConnections < 6 {
		return false
	}

	return true
}

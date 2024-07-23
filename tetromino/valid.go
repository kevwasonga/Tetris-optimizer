package tetromino

import (
	"strings"
)

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
		for j, char := range line {
			if char != '.' && char != '#' {
				return false
			}
			if char == '#' {
				countHashes++
				connections := 0
				for _, dir := range directions {
					ni, nj := i+dir.x, j+dir.y
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(line) && lines[ni][nj] == '#' {
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

	if countHashes != 4 || totalConnections < 6 {
		return false
		// os.Exit(0)
	}

	return true
}

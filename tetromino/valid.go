package tetromino

import "strings"

// validate tetrominoes
func IsValidTetromino(tetromino string) bool {
	lines := strings.Split(tetromino, "\n")
	// Ensure there are exactly 4 lines
	if len(lines) != 4 {
		return false
	}
	// Ensure each line has the correct length
	for _, line := range lines {
		if len(line) != 4 {
			return false
		}
	}
	return true
}

package tetromino

import "strings"

// convertStringToRuneArray converts a string representation of a tetromino to a 2D rune array
func ConvertStringToRuneArray(tetrominoStr string) [][]rune {
	lines := strings.Split(tetrominoStr, "\n")
	array := make([][]rune, len(lines))
	for i, line := range lines {
		array[i] = []rune(line)
	}
	return array
}

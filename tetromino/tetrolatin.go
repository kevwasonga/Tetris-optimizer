package tetromino

import "strings"

func ReplaceChars(tetromino string, key int) string {
	var result strings.Builder
	letter := rune('A' + key - 1) // Calculate the corresponding letter

	for _, char := range tetromino {
		if char == '#' {
			result.WriteRune(letter) // Replace '#' with the corresponding letter
		} else {
			result.WriteRune(char) // Keep other characters unchanged
		}
	}
	return result.String()
}

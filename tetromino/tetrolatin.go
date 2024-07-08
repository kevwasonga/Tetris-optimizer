package tetromino

import "strings"

func ReplaceChars(tetromino string, key int) string {
	var result strings.Builder
	for _, char := range tetromino {
		switch key {
		case 1:
			if char == '#' {
				result.WriteRune('A')
			} else {
				result.WriteRune(char)
			}
		case 2:
			if char == '#' {
				result.WriteRune('B')
			} else {
				result.WriteRune(char)
			}
		case 3:
			if char == '#' {
				result.WriteRune('C')
			} else {
				result.WriteRune(char)
			}
		case 4:
			if char == '#' {
				result.WriteRune('D')
			} else {
				result.WriteRune(char)
			}
		case 5:
			if char == '#' {
				result.WriteRune('E')
			} else {
				result.WriteRune(char)
			}
		case 6:
			if char == '#' {
				result.WriteRune('F')
			} else {
				result.WriteRune(char)
			}
		case 7:
			if char == '#' {
				result.WriteRune('G')
			} else {
				result.WriteRune(char)
			}
		case 8:
			if char == '#' {
				result.WriteRune('H')
			} else {
				result.WriteRune(char)
			}
		// Add more cases as needed for other keys
		default:
			result.WriteRune(char)
		}
	}
	return result.String()
}

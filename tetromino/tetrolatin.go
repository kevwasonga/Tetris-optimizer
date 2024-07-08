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
		case 9:
			if char == '#' {
				result.WriteRune('I')
			} else {
				result.WriteRune(char)
			}
		case 10:
			if char == '#' {
				result.WriteRune('J')
			} else {
				result.WriteRune(char)
			}
		case 11:
			if char == '#' {
				result.WriteRune('K')
			} else {
				result.WriteRune(char)
			}
		case 12:
			if char == '#' {
				result.WriteRune('L')
			} else {
				result.WriteRune(char)
			}
		case 13:
			if char == '#' {
				result.WriteRune('M')
			} else {
				result.WriteRune(char)
			}
		case 14:
			if char == '#' {
				result.WriteRune('N')
			} else {
				result.WriteRune(char)
			}
		case 15:
			if char == '#' {
				result.WriteRune('O')
			} else {
				result.WriteRune(char)
			}
		case 16:
			if char == '#' {
				result.WriteRune('P')
			} else {
				result.WriteRune(char)
			}
		case 17:
			if char == '#' {
				result.WriteRune('Q')
			} else {
				result.WriteRune(char)
			}
		case 18:
			if char == '#' {
				result.WriteRune('R')
			} else {
				result.WriteRune(char)
			}
		case 19:
			if char == '#' {
				result.WriteRune('S')
			} else {
				result.WriteRune(char)
			}
		case 20:
			if char == '#' {
				result.WriteRune('T')
			} else {
				result.WriteRune(char)
			}
		default:
			result.WriteRune(char)
		}
	}
	return result.String()
}

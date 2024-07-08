package tetromino

import (
	"bufio"
	"fmt"
	"os"
)

// map the banner file containing the tetrominoes
func Loadbanner(fileName string) (map[int]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening file %s: %v", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int]string)
	key := 1
	linecount := 0
	var random string

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			if linecount < 3 {
				random += line + "\n"
			} else {
				random += line
			}
			linecount++
		}
		if linecount == 4 {
			bannerMap[key] = random
			key++
			linecount = 0
			random = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file %s: %v", fileName, err)
	}

	return bannerMap, nil
}

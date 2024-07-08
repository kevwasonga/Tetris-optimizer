package main

import (
	"fmt"

	"tetris/tetromino"
)

func main() {
	bannerMapmap := tetromino.Loadbanner
	tetrominoes, err := bannerMapmap("sample.txt")
	if err != nil {
		fmt.Println("Error loading tetrominoes:", err)
		return
	}

	valid := tetromino.IsValidTetromino

	validTetrominoes := make(map[int]string)
	for key, tetrominoe := range tetrominoes {
		if valid(tetrominoe) {
			validTetrominoes[key] = tetromino.ReplaceChars(tetrominoe, key)
		} else {
			fmt.Printf("Tetromino %d is invalid:\n%s\n", key, tetrominoe)
		}
	}

	if len(validTetrominoes) == 0 {
		fmt.Println("No valid tetrominoes found.")
		return
	}

	print := tetromino.AssembleTetrominoes

	square := print(validTetrominoes)
	fmt.Println("Assembled Tetrominoes into Square:")
	printSquare(square)
}

func printSquare(square [][]rune) {
	for _, row := range square {
		fmt.Println(string(row))
	}
}

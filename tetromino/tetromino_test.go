package tetromino

import (
	"testing"
)

func TestIsValidTetromino(t *testing.T) {
	validTetromino := "....\n.##.\n.##.\n...."
	invalidTetromino := "....\n.###.\n.##.\n...."

	if !IsValidTetromino(validTetromino) {
		t.Errorf("Expected valid tetromino to be valid")
	}

	if IsValidTetromino(invalidTetromino) {
		t.Errorf("Expected invalid tetromino to be invalid")
	}
}

func TestConvertStringToRuneArray(t *testing.T) {
	input := "....\n.##.\n.##.\n...."
	expected := [][]rune{
		{'.', '.', '.', '.'},
		{'.', '#', '#', '.'},
		{'.', '#', '#', '.'},
		{'.', '.', '.', '.'},
	}

	result := ConvertStringToRuneArray(input)

	// Check if the result matches the expected output
	for i, row := range result {
		for j, value := range row {
			if value != expected[i][j] {
				t.Errorf("At position (%d, %d), expected '%c', but got '%c'", i, j, expected[i][j], value)
			}
		}
	}
}

func TestCalculateMinimumSquareSize(t *testing.T) {
	tests := []struct {
		numTetrominoes int
		expectedSize   int
	}{
		{1, 2},
		{2, 3},
		{4, 4},
		{10, 7},
	}

	for _, test := range tests {
		result := CalculateMinimumSquareSize(test.numTetrominoes)
		if result != test.expectedSize {
			t.Errorf("For %d tetrominoes, expected board size %d, but got %d", test.numTetrominoes, test.expectedSize, result)
		}
	}
}

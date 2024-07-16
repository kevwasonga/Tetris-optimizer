package tetromino

import (
	"os"
	"testing"
)

// TestLoadBanner tests loading a banner file and ensures tetrominoes are loaded correctly.
func TestLoadBanner(t *testing.T) {
	testFile := "test_banner.txt"
	bannerContent := `#..#
###.
....
....
`
	err := os.WriteFile(testFile, []byte(bannerContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	bannerMap, err := Loadbanner(testFile)
	if err != nil {
		t.Fatalf("Loadbanner returned an error: %v", err)
	}

	if len(bannerMap) != 1 {
		t.Errorf("Expected 1 tetromino, got %d", len(bannerMap))
	}
}

// TestIsValidTetromino tests the validity of tetrominoes.
func TestIsValidTetromino(t *testing.T) {
	validTetromino := "#..#\n###.\n....\n...."
	if IsValidTetromino(validTetromino) {
		t.Errorf("Expected valid tetromino to be valid")
	}

	invalidTetromino := "#...\n#.#.\n....\n...."
	if IsValidTetromino(invalidTetromino) {
		t.Errorf("Expected invalid tetromino to be invalid")
	}
}

// TestAssembleTetrominoes tests assembling tetrominoes onto the board.
func TestAssembleTetrominoes(t *testing.T) {
	tetrominoes := map[rune]string{
		'A': "#..#\n###.\n....\n....",
		'B': ".#..\n###.\n....\n....",
	}

	board := AssembleTetrominoes(tetrominoes)
	if board == nil {
		t.Error("AssembleTetrominoes returned nil board")
	}

	// Check that tetrominoes A and B are placed on the board
	countA, countB := 0, 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 'A' {
				countA++
			} else if cell == 'B' {
				countB++
			}
		}
	}
	if countA == 0 {
		t.Error("Tetromino A not placed on the board")
	}
	if countB == 0 {
		t.Error("Tetromino B not placed on the board")
	}
}

// Helper function to capture output
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = w

	f()

	w.Close()
	out, _ := os.ReadFile(r.Name())
	return string(out)
}

// TestRemoveTetromino tests removing a tetromino from the board.
func TestRemoveTetromino(t *testing.T) {
	board := initializeBoard(4)
	tetromino := "#..#\n###.\n....\n...."

	PutTetromino(tetromino, board, 0, 0, 'A')
	RemoveTetromino(tetromino, board, 0, 0)

	// Check if the board is reset to '.'
	for _, row := range board {
		for _, cell := range row {
			if cell != '.' {
				t.Error("Failed to remove tetromino, found non-empty cell")
			}
		}
	}
}

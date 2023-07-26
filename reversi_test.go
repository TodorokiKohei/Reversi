package main

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if len(board.tokens) != Size {
		t.Errorf("Expected board to have size %d, got %d", Size, len(board.tokens))
	}
	for i := 0; i < Size; i++ {
		if len(board.tokens[i]) != Size {
			t.Errorf("Expected board row to have size %d, got %d", Size, len(board.tokens[i]))
		}
	}
}

func TestIsValidMove(t *testing.T) {
	board := NewBoard()
	if isValidMove(board, 0, 0, Black) {
		t.Errorf("Expected move to be invalid, but it was valid")
	}
	if !isValidMove(board, 2, 3, Black) {
		t.Errorf("Expected move to be valid, but it was invalid")
	}
}

func TestMakeMove(t *testing.T) {
	board := NewBoard()
	makeMove(board, 2, 3, Black)
	if board.tokens[3][3] != Black {
		t.Errorf("Expected token at position (3,3) to be Black, got %d", board.tokens[3][3])
	}
}

func TestEvaluateBoard(t *testing.T) {
	board := NewBoard()
	if evaluateBoard(board, Black) != 0 {
		t.Errorf("Expected initial board evaluation for Black to be 0, got %f", evaluateBoard(board, Black))
	}
}

func TestParseInput(t *testing.T) {
	_, _, err := parseInput("z9")
	if err == nil {
		t.Errorf("Expected error for out of range input")
	}

	_, _, err = parseInput("b0")
	if err == nil {
		t.Errorf("Expected error for out of range input")
	}

	x, y, err := parseInput("h8")
	if err != nil || x != 7 || y != 7 {
		t.Errorf("Expected x=7, y=7, got x=%d, y=%d", x, y)
	}
}

func TestGetBestMove(t *testing.T) {
	b := NewBoard()
	move := getBestMove(b, Black)
	expectedMove := []int{3, 2}
	if !reflect.DeepEqual(move, expectedMove) {
		t.Errorf("Expected move %v, got %v", expectedMove, move)
	}
}

// TestIsGameOver tests the isGameOver function
func TestIsGameOver(t *testing.T) {
	b := NewBoard()
	if isGameOver(b) {
		t.Errorf("Expected game to be not over, got over")
	}

	// make a board where the game is over
	for i := range b.tokens {
		for j := range b.tokens[i] {
			b.tokens[i][j] = Black
		}
	}
	if !isGameOver(b) {
		t.Errorf("Expected game to be over, got not over")
	}
}

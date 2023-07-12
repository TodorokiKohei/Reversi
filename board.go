package main

import (
	//"bufio"
	"fmt"
	//"os"
	//	"strconv"
	//	"strings"
)

// define constant number
const (
	empty = 0
	black = 1
	white = -1
	size  = 8
)

type Board struct {
	tokens [][]int
}

// Create a new board class
func NewBoard() *Board {
	return &Board{
		tokens: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

}

/*
// Initialization of chessboard
func (b *Board) Init() {
	b[3][3] = -1
	b[4][4] = -1
	b[3][4] = 1
	b[4][3] = 1
}
*/

// put a specific chess on chessboard
func (b *Board) Put(x, y int, u string) {
	if u == "b" {
		b.tokens[y][x] = 1
	} else if u == "w" {
		b.tokens[y][x] = -1
	}
}

// get info of chess from a position
func (b *Board) Get(x, y int) string {
	if b.tokens[y][x] == 1 {
		return "b"
	} else if b.tokens[y][x] == -1 {
		return "w"
	}
	return "n"
}

// Display function of chessboard
func (b *Board) Print() {
	fmt.Println("  1 2 3 4 5 6 7 8")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 8; j++ {
			switch b.Get(j, i) {
			case "b":
				fmt.Print("b ")
			case "w":
				fmt.Print("w ")
			case "n":
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

// get player side
func getPlayerString(player int) string {
	switch player {
	case 1:
		return "b"
	case -1:
		return "w"
	default:
		return "n"
	}
}

//Determine the winner and count the number of Othello pieces.
func countPieces(board *Board) (blackCount, whiteCount int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch board.Get(j, i) {
			case "b":
				blackCount++
			case "w":
				whiteCount++
			}
		}
	}
	return blackCount, whiteCount
}

func main() {
	board := NewBoard()
	//	reader := bufio.NewReader(os.Stdin)
	//	player := 1
	board.Print()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// define constant number
const (
	empty = 0
	black = 1
	white = -1
	size  = 8
)

func getPlayerName(player int) string {
	switch player {
	case black:
		return "Black"
	case white:
		return "White"
	default:
		return ""
	}
}

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

// Initialization of chessboard
func (b *Board) Init() {
	b.tokens[3][3] = white
	b.tokens[4][4] = white
	b.tokens[3][4] = black
	b.tokens[4][3] = black
}

// put a specific chess on chessboard
func (b *Board) Put(x, y int, u int) {
	if u == black {
		b.tokens[y][x] = 1
	} else if u == white {
		b.tokens[y][x] = -1
	}
}

// get info of chess from a position
func (b *Board) Get(x, y int) int {
	if b.tokens[y][x] == 1 {
		return black
	} else if b.tokens[y][x] == -1 {
		return white
	}
	return 0
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
	board.Init()

	reader := bufio.NewReader(os.Stdin)
	player := black
	for {
		board.Print()

		fmt.Printf("Player %s's turn. Enter move (row col): ", getPlayerName(player))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		move := strings.Split(input, " ")
		if len(move) != 2 {
			fmt.Println("Invalid move. Enter row and col separated by a space.")
			continue
		}

		row, err := strconv.Atoi(move[0])
		if err != nil {
			fmt.Println("Invalid move. Row must be a number.")
			continue
		}

		col, err := strconv.Atoi(move[1])
		if err != nil {
			fmt.Println("Invalid move. Column must be a number.")
			continue
		}
		xInput := col - 1
		yInput := row - 1

		board.Put(xInput, yInput, player)

		player = -player
	}
}

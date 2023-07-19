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
func (b *Board) Put(x, y, u int) {
	if u == black {
		b.tokens[y][x] = 1
	} else if u == white {
		b.tokens[y][x] = -1
	}
}

// get info of chess from a position
func (b *Board) Get(x, y int) string {
	if b.tokens[y][x] == black {
		return "b"
	} else if b.tokens[y][x] == white {
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

func (b *Board) CanPut(x, y, u int) bool {
	var other, mycolor string
	if u == black {
		mycolor = "b"
		other = "w"
	} else {
		mycolor = "w"
		other = "b"
	}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// 違う色判定 && 範囲外でないか
			if x+i < 0 || x+i >= size || y+j < 0 || y+j >= size {
				continue
			}
			if b.Get(x+i, y+j) != other {
				continue
			}

			// 違う色であれば、その方向を進めて反転できる条件にマッチしているか
			for s := 2; s <= size; s++ {
				if x+i*s >= 0 &&
					x+i*s < size &&
					y+j*s >= 0 &&
					y+j*s < size {
					if b.Get(x+i*s, y+j*s) == "n" {
						break
					}
					if b.Get(x+i*s, y+j*s) == mycolor {
						return true
					}

				}
			}
		}
	}
	return false
}

func (b *Board) reversal(x, y, u int) {
	var other, mycolor string
	if u == black {
		mycolor = "b"
		other = "w"
	} else {
		mycolor = "w"
		other = "b"
	}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// 違う色判定 && 範囲外でないか
			if x+i < 0 || x+i >= size || y+j < 0 || y+j >= size {
				continue
			}
			if b.Get(x+i, y+j) != other {
				continue
			}

			// 違う色であれば、その方向を進めて反転できる条件にマッチしているか
			for s := 2; s <= size; s++ {
				if x+i*s >= 0 &&
					x+i*s < size &&
					y+j*s >= 0 &&
					y+j*s < size {
					if b.Get(x+i*s, y+j*s) == "n" {
						break
					}
					if b.Get(x+i*s, y+j*s) == mycolor {
						b.Put(x, y, u)
						for n := 1; n < s; n++ {
							b.Put(x+i*n, y+j*n, u)
						}
						break
					}
				}
			}
		}
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

func getPlayerName(u int) string {
	if u == black {
		return "black"
	} else if u == white {
		return "white"
	}
	return ""
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

		if !board.CanPut(xInput, yInput, player) {
			fmt.Println("Invalid move. Can't put.")
			continue
		}
		board.Put(xInput, yInput, player)
		board.reversal(xInput, yInput, player)

		player = -player
	}
}

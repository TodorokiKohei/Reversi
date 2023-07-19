package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	tokens [][]int
}

const (
	Empty = 0
	Black = 1
	White = 2
	Size  = 8
)

func NewBoard() *Board {
	return &Board{
		tokens: [][]int{
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, White, Black, Empty, Empty, Empty},
			{Empty, Empty, Empty, Black, White, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		},
	}
}

func (b *Board) Print() {
	fmt.Println("  a b c d e f g h")
	for i := 0; i < Size; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < Size; j++ {
			switch b.tokens[i][j] {
			case Empty:
				fmt.Print(". ")
			case Black:
				fmt.Print("B ")
			case White:
				fmt.Print("W ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) Put(x, y, player int) {
	b.tokens[y][x] = player
}

func parseInput(move string) (int, int, error) {
	if len(move) != 2 {
		return -1, -1, fmt.Errorf("Invalid move format")
	}
	col := int(move[0] - 'a')
	row, err := strconv.Atoi(string(move[1]))
	if err != nil || row < 1 || row > Size {
		return -1, -1, fmt.Errorf("Invalid move format")
	}
	return col, row - 1, nil
}

func isValidMove(board *Board, x, y, player int) bool {
	if x < 0 || x >= Size || y < 0 || y >= Size || board.tokens[y][x] != Empty {
		return false
	}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		tx, ty := x+dx[i], y+dy[i]
		if tx >= 0 && tx < Size && ty >= 0 && ty < Size && board.tokens[ty][tx] == 3-player {
			for {
				tx, ty = tx+dx[i], ty+dy[i]
				if tx < 0 || tx >= Size || ty < 0 || ty >= Size {
					break
				}
				if board.tokens[ty][tx] == player {
					return true
				}
				if board.tokens[ty][tx] == Empty {
					break
				}
			}
		}
	}

	return false
}

func makeMove(board *Board, x, y, player int) {
	board.Put(x, y, player)

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		tx, ty := x+dx[i], y+dy[i]
		flip := false
		for {
			if tx < 0 || tx >= Size || ty < 0 || ty >= Size {
				break
			}
			if board.tokens[ty][tx] == Empty {
				break
			}
			if board.tokens[ty][tx] == player {
				flip = true
				break
			}
			tx, ty = tx+dx[i], ty+dy[i]
		}
		if flip {
			tx, ty = x+dx[i], y+dy[i]
			for board.tokens[ty][tx] != player {
				board.Put(tx, ty, player)
				tx, ty = tx+dx[i], ty+dy[i]
			}
		}
	}
}

func getPlayerName(player int) string {
	switch player {
	case Black:
		return "Black"
	case White:
		return "White"
	default:
		return ""
	}
}

func main() {
	board := NewBoard()

	reader := bufio.NewReader(os.Stdin)

	player := Black
	for {
		board.Print()

		if player == Black {
			fmt.Printf("Player %s's turn. Enter move (e.g., a1): ", getPlayerName(player))
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			input = strings.TrimSpace(input)
			col, row, err := parseInput(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !isValidMove(board, col, row, player) {
				fmt.Println("Invalid move. Try again.")
				continue
			}
			makeMove(board, col, row, player)
		} else {
			fmt.Printf("AI %s's turn...\n", getPlayerName(player))
			validMoves := findValidMoves(board, player)
			if len(validMoves) == 0 {
				fmt.Println("AI has no valid move. Skipping turn.")
			} else {
				bestMove := getBestMove(board, validMoves, player)
				makeMove(board, bestMove[0], bestMove[1], player)
				fmt.Printf("AI %s made a move at %s\n", getPlayerName(player), formatMove(bestMove))
			}
		}

		player = 3 - player
	}
}

func findValidMoves(board *Board, player int) [][]int {
	validMoves := [][]int{}
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			if isValidMove(board, x, y, player) {
				validMoves = append(validMoves, []int{x, y})
			}
		}
	}
	return validMoves
}

func getBestMove(board *Board, moves [][]int, player int) []int {
	bestMove := moves[0]
	maxFlips := 0
	for _, move := range moves {
		flips := countFlips(board, move[0], move[1], player)
		if flips > maxFlips {
			maxFlips = flips
			bestMove = move
		}
	}
	return bestMove
}

func countFlips(board *Board, x, y, player int) int {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	flips := 0
	for i := 0; i < 8; i++ {
		tx, ty := x+dx[i], y+dy[i]
		count := 0
		for {
			if tx < 0 || tx >= Size || ty < 0 || ty >= Size {
				break
			}
			if board.tokens[ty][tx] == player {
				flips += count
				break
			}
			if board.tokens[ty][tx] == Empty {
				break
			}
			tx, ty = tx+dx[i], ty+dy[i]
			count++
		}
	}
	return flips
}

func formatMove(move []int) string {
	col := string('a' + move[0])
	row := strconv.Itoa(move[1] + 1)
	return col + row
}

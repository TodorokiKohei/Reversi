package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
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
				fmt.Print("● ")
			case White:
				fmt.Print("o ")
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

func randomMove(board *Board, player int) []int {
	validMoves := findValidMoves(board, player)
	if len(validMoves) == 0 {
		return nil
	}
	return validMoves[rand.Intn(len(validMoves))]
}

func getBestMove(board *Board, player int) []int {
	validMoves := findValidMoves(board, player)
	if len(validMoves) == 0 {
		return nil
	}

	bestMove := validMoves[0]
	bestScore := math.Inf(-1)

	for _, move := range validMoves {
		newBoard := copyBoard(board)
		makeMove(newBoard, move[0], move[1], player)
		score := evaluateBoard(newBoard, player)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

func minimax(board *Board, player, depth int) []int {
	validMoves := findValidMoves(board, player)
	if len(validMoves) == 0 || depth == 0 {
		return nil
	}

	bestMove := validMoves[0]
	bestScore := math.Inf(-1)

	for _, move := range validMoves {
		newBoard := copyBoard(board)
		makeMove(newBoard, move[0], move[1], player)
		score := -negamax(newBoard, 3-player, depth-1)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

func negamax(board *Board, player, depth int) float64 {
	validMoves := findValidMoves(board, player)
	if len(validMoves) == 0 || depth == 0 {
		return evaluateBoard(board, player)
	}

	bestScore := math.Inf(-1)

	for _, move := range validMoves {
		newBoard := copyBoard(board)
		makeMove(newBoard, move[0], move[1], player)
		score := -negamax(newBoard, 3-player, depth-1)
		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
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

func evaluateBoard(board *Board, player int) float64 {
	blackCount, whiteCount := countPieces(board)
	if player == Black {
		return float64(blackCount - whiteCount)
	}
	return float64(whiteCount - blackCount)
}

func countPieces(board *Board) (blackCount, whiteCount int) {
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			switch board.tokens[y][x] {
			case Black:
				blackCount++
			case White:
				whiteCount++
			}
		}
	}
	return blackCount, whiteCount
}

func copyBoard(board *Board) *Board {
	newBoard := NewBoard()
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			newBoard.tokens[y][x] = board.tokens[y][x]
		}
	}
	return newBoard
}

func formatMove(move []int) string {
	col := string('a' + move[0])
	row := strconv.Itoa(move[1] + 1)
	return col + row
}

func isGameOver(board *Board) bool {
	return len(findValidMoves(board, Black)) == 0 && len(findValidMoves(board, White)) == 0
}

func getWinner(board *Board) int {
	blackCount, whiteCount := countPieces(board)
	if blackCount > whiteCount {
		fmt.Println("blackCount are ", blackCount)
		fmt.Println("whiteCount are ", whiteCount)
		return Black
	} else if whiteCount > blackCount {
		fmt.Println("blackCount are ", blackCount)
		fmt.Println("whiteCount are ", whiteCount)
		return White
	}
	return Empty
}

func main() {
	board := NewBoard()
	reader := bufio.NewReader(os.Stdin)
	var gameMode int
	var err error
	for {
		fmt.Println("Welcome to Othello!")
		fmt.Println("Select game mode:")
		fmt.Println("1. Player vs Player")
		fmt.Println("2. Player vs AI")
		gameMode, err = readGameMode(reader)
		if (gameMode != 1) && (gameMode != 2) {
			fmt.Println("Invalid game mode. please input again")
		} else {
			break
		}
	}

	aiLevel := 0
	for {
		if gameMode == 2 {
			fmt.Println("Select AI level:")
			fmt.Println("1. Beginner (Random Move)")
			fmt.Println("2. Intermediate (Greedy Move)")
			fmt.Println("3. Advanced (Minimax Search)")
			aiLevel, err = readAILevel(reader)
		}
		if aiLevel != 1 && aiLevel != 2 && aiLevel != 3 {
			fmt.Println("Invalid game mode. please input again")
		} else {
			break
		}
	}
	_ = err

	player := Black

	for {
		board.Print()

		if !hasValidMove(board, player) {
			fmt.Printf("Player %s has no valid move. Skipping turn.\n", getPlayerName(player))
			player = 3 - player
			continue
		}

		if player == Black || gameMode == 1 {
			fmt.Printf("Player %s's turn. Enter move (e.g., a1): ", getPlayerName(player))
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			input = strings.TrimSpace(input)
			if input == "exit" {
				fmt.Println("Game exited.")
				break
			}
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
			var bestMove []int
			switch aiLevel {
			case 1:
				bestMove = randomMove(board, player)
			case 2:
				bestMove = getBestMove(board, player)
			case 3:
				bestMove = minimax(board, player, 4) //The depth of the game tree, which determines the optimal solution for the AI up to the depth step
			default:
				bestMove = randomMove(board, player)
			}

			makeMove(board, bestMove[0], bestMove[1], player)
			fmt.Printf("AI %s made a move at %s\n", getPlayerName(player), formatMove(bestMove))
		}

		if isGameOver(board) {
			board.Print()
			winner := getWinner(board)
			switch winner {
			case Black:
				fmt.Println("Game over. Black wins!")
			case White:
				fmt.Println("Game over. White wins!")
			default:
				fmt.Println("Game over. It's a draw!")
			}
			break
		}

		player = 3 - player
	}
}

func readAILevel(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	aiLevel, err := strconv.Atoi(input)
	if err != nil || (aiLevel != 1 && aiLevel != 2 && aiLevel != 3) {
		return 0, fmt.Errorf("Invalid AI level")
	}
	return aiLevel, nil
}

func readGameMode(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	gameMode, err := strconv.Atoi(input)
	if err != nil || (gameMode < 1 || gameMode > 4) {
		return 0, fmt.Errorf("Invalid game mode")
	}
	return gameMode, nil
}

func hasValidMove(board *Board, player int) bool {
	validMoves := findValidMoves(board, player)
	return len(validMoves) > 0
}

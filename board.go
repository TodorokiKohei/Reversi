package Reversi

type Board struct {
	tokens [][]int
}

func NewBoard() *Board{
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

func (b *Board) Put(x, y int, u string) {
    if u == "b" {
        b.tokens[y][x] = 1   
    } else if u == "w" {
        b.tokens[y][x] = -1
    }
}


func (b *Board) Get(x, y int) string{
    if b.tokens[y][x] == 1 {
        return "b"
    } else if b.tokens[y][x] == -1 {
        return "w"
    }
    return "n"
}
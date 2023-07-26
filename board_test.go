package main

import "testing"
import "reflect"

func TestPutToken01(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, white)
	if b.Get(1, 1) != "w" {
		t.Errorf("Test fail expected: %s, result: %s\n", "w", b.Get(1, 1))
	}

	b.Put(2, 2, black)
	if b.Get(2, 2) != "b" {
		t.Errorf("Test fail expected: %s, result: %s\n", "b", b.Get(2, 2))
	}
}

func TestCanPut01(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, white)
	b.Put(1, 2, black)
	if b.CanPut(1, 3, white) != true {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.CanPut(1, 3, white))
	}
}

func TestCanPut02(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, white)
	b.Put(1, 2, black)
	if b.CanPut(1, 5, white) != false {
		t.Errorf("Test fail expected: %t, result: %t\n", false, b.CanPut(1, 5, white))
	}
}

func TestCanPut03(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, white)
	b.Put(2, 2, white)
	b.Put(0, 0, black)
	if b.CanPut(3, 3, black) != true {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.CanPut(3, 3, black))
	}
}

func TestRereversal(t *testing.T) {
	b := NewBoard()
	eb := NewBoard()
	b.Put(1, 1, white)
	b.Put(2, 2, white)
	b.Put(0, 0, black)
	b.reversal(3, 3, black)

	eb.Put(1, 1, black)
	eb.Put(2, 2, black)
	eb.Put(0, 0, black)
	eb.Put(3, 3, black)
	b.Print()
	eb.Print()

	if !reflect.DeepEqual(b.tokens, eb.tokens) {
		t.Errorf("Test fail expected")
	}
}

func TestCanPlay01(t *testing.T) {
	b := NewBoard()
	b.Init()
	if b.CanPlay(black) != true {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.CanPlay(black))
	}
}

func TestCanPlay02(t *testing.T) {
	b := NewBoard()
	if b.CanPlay(black) != false {
		t.Errorf("Test fail expected: %t, result: %t\n", false, b.CanPlay(black))
	}
}

func TestCanPlay03(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, white)
	b.Put(1, 2, black)
	if b.CanPlay(white) != true {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.CanPlay(white))
	}
}

func TestIsGameOver01(t *testing.T) {
	b := NewBoard()
	b.Init()
	if b.IsGameOver() != false {
		t.Errorf("Test fail expected: %t, result: %t\n", false, b.IsGameOver())
	}
}

func TestIsGameOver02(t *testing.T) {
	b := NewBoard()
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			b.Put(x, y, black)
		}
	}
	if b.IsGameOver() != true {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.IsGameOver())
	}
}

func TestIsGameOver03(t *testing.T) {
	b := NewBoard()
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if (x+y)%2 == 0 {
				b.Put(x, y, black)
			} else {
				b.Put(x, y, white)
			}
		}
	}
	if b.IsGameOver() != true {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.IsGameOver())
	}
}

func TestGetWinner01(t *testing.T) {
	b := NewBoard()
	b.Put(0, 0, black)
	b.Put(1, 0, black)
	b.Put(0, 1, white)
	if b.GetWinner() != "Black" {
		t.Errorf("Test fail expected: %s, result: %s\n", "Black", b.GetWinner())
	}
}

func TestGetWinner02(t *testing.T) {
	b := NewBoard()
	b.Put(0, 0, white)
	b.Put(1, 0, black)
	b.Put(0, 1, white)
	if b.GetWinner() != "White" {
		t.Errorf("Test fail expected: %s, result: %s\n", "White", b.GetWinner())
	}
}

func TestGetWinner03(t *testing.T) {
	b := NewBoard()
	b.Put(0, 0, black)
	b.Put(1, 0, white)
	if b.GetWinner() != "Draw" {
		t.Errorf("Test fail expected: %s, result: %s\n", "Draw", b.GetWinner())
	}
}

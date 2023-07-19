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

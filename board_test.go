package main

import "testing"

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

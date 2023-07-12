package main

import "testing"

func TestPutToken01(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, "w")
	if b.Get(1, 1) != "w" {
		t.Errorf("Test fail expected: %s, result: %s\n", "w", b.Get(1, 1))
	}

	b.Put(2, 2, "b")
	if b.Get(2, 2) != "b" {
		t.Errorf("Test fail expected: %s, result: %s\n", "b", b.Get(2, 2))
	}
}

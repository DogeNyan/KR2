package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 7
	a, b := 3, 4
	got := a + b
	if got != expected {
		t.Errorf("add(%d,%d) = %d; need %d", a, b, got, expected)
	}
}

package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "oven" {
		t.Errorf("expected: even, actual: %s", result)
	}
}

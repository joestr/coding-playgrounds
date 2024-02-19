package main

import "testing"

func TestOnePlusOne(t *testing.T) {
	var result = onePlusOne()
	if result != 2 {
		t.Fatalf("1 + 1 MUST BE 2!!")
	}
}

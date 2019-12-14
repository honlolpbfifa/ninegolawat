package main

import "testing"

func TestPlus(T *testing.T) {
	result := Plus(1, 5)

	if result != 6 {
		t.Error("Expected 6, got : ", result)
	}

}

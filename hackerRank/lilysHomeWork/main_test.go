package main

import "testing"

func TestSwapDesc(t *testing.T) {
	arr := []int32{3, 4, 2, 5, 1}
	swap := lilysHomework(arr)
	answer := 2
	if swap != int32(answer) {
		t.Errorf("Expected %d , got %d",answer,swap)
	}
}
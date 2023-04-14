package main

import (
	"testing"
)

func compareArrays(a *[7]int, b *[7]int) bool {
	for i,j := range a{
		if j != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	toTest := [7]int{4,5,1,7,3,9,5}
	got := bubbleSort(&toTest)
	want := [7]int{1,3,4,5,5,7,9}
	

	if !compareArrays(&got, &want) {
		t.Errorf("got %v want %v", got, want)
	}
}
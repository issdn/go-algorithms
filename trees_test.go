package main

import "testing"

var arr = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

func compareArrays(arr1, arr2 []int) bool {
	for i, j := range arr1 {
		if j != arr2[i] {
			return false
		}
	}
	return true
}

func ExampleTree() *Node {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}
	return root
}

func TestBinarySearch(t *testing.T) {
	got := BinarySearch(arr, 10)
	want := 4

	if got != want {
		t.Errorf("Wanted %v, got %v.", want, got)
	}
}

func TestBFS(t *testing.T) {
	got := BFS(ExampleTree())
	want := []int{1, 2, 3, 4, 5, 6, 7}
	equal := compareArrays(got, want)
	if !equal {
		t.Errorf("Wanted %v, got %v.", want, got)
	}
}

package main

import (
	"errors"
	"fmt"
	"testing"
)

func compareEqualLenArrays(a, b []uint16) (bool, error) {
	if len(a) != len(b) {
		return false, errors.New("Arrays are not of equal length")
	}
	for i, j := range a {
		if j != b[i] {
			return false, errors.New("Arrays are not of equal")
		}
	}
	return true, nil
}

type SortTestError struct {
	TestCase []uint16
	Got      []uint16
	Want     []uint16
}

func (e *SortTestError) FormatTestErrorMessage() string {
	return fmt.Sprintf("\nTest Case:%v\nGot: %v\nWant: %v", e.TestCase, e.Got, e.Want)
}

func TestBubbleSort(t *testing.T) {
	arrToTest := [...]uint16{4, 5, 1, 7, 3, 9, 5}
	arrToSort := make([]uint16, len(arrToTest))
	copy(arrToSort, arrToTest[:])
	got := bubbleSort(arrToTest[:])
	want := [...]uint16{1, 3, 4, 5, 5, 7, 9}

	isEqual, error := compareEqualLenArrays(got, want[:])
	if error != nil {
		t.Errorf("Error: %v", error)
	}

	if !isEqual {
		t.Error((&SortTestError{arrToTest[:], got, want[:]}).FormatTestErrorMessage())
	}
}

func TestQuicksort(t *testing.T) {
	want := [...]uint16{1, 3, 4, 5, 5, 7, 9}
	arrToTest := [...]uint16{4, 5, 1, 7, 3, 9, 5}
	arrToSort := make([]uint16, len(arrToTest))
	copy(arrToSort, arrToTest[:])
	got := quicksort(arrToSort, 0, uint16(len(arrToTest)-1))

	isEqual, error := compareEqualLenArrays(got, want[:])
	if error != nil {
		t.Errorf("Error: %v", error)
	}

	if !isEqual {
		t.Error((&SortTestError{arrToTest[:], got, want[:]}).FormatTestErrorMessage())
	}
}

package main

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	target := 3
	sort.Ints(arr)
	index := BinarySearch(arr, target)
	if index != 1 {
		t.Errorf("Expected index %d, but got %d", 1, index)
	}

	target = 2
	index = BinarySearch(arr, target)
	if index != -1 {
		t.Errorf("Expected index %d, but got %d", -1, index)
	}
}

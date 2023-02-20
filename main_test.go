package main

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 5
	sort.Ints(arr)
	index := BinarySearch(arr, target)
	if index != 2 {
		t.Errorf("Expected index %d, but got %d", 2, index)
	}

	target = 6
	index = BinarySearch(arr, target)
	if index != -1 {
		t.Errorf("Expected index %d, but got %d", -1, index)
	}
}

func TestSumTwoInt(t *testing.T) {
	sum := SumTwoInt(1, 3)
	if sum != 4 {
		t.Errorf("Expected sum %v, but got %v", 4, sum)
	}

	sum = SumTwoInt(3, 7)
	if sum != 10 {
		t.Errorf("Expected sum %v, but got %v", 10, sum)
	}

}

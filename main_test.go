package main

import (
	"reflect"
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

func TestTwoSum(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10, 12}
	target := 22
	anwser := TwoSum(nums, target)
	expect := []int{4, 5}
	if reflect.DeepEqual(anwser, expect) {
		t.Errorf("Expected anwser %v,but got %v", expect, anwser)
	}

	nums = []int{3, 2, 4}
	target = 6
	anwser = TwoSum(nums, target)
	expect = []int{1, 2}
	if reflect.DeepEqual(anwser, expect) {
		t.Errorf("Expected anwser %v,but got %v", expect, anwser)
	}
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	anwser := SearchInsert(nums, target)
	if anwser != 2 {
		t.Errorf("Expeced anwser %v, but got %v", 2, anwser)
	}

	target = 7
	anwser = SearchInsert(nums, target)
	if anwser != 4 {
		t.Errorf("Expeced anwser %v, but got %v", 4, anwser)
	}

	nums = []int{1, 3}
	target = 1
	anwser = SearchInsert(nums, target)
	if anwser != 0 {
		t.Errorf("Expeced anwser %v, but got %v", 0, anwser)
	}
}

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	if anwser := IsValid(s); !anwser {
		t.Errorf("Expeced anwser true, but got %v", anwser)
	}

	s = ")("
	if anwser := IsValid(s); anwser {
		t.Errorf("Expeced anwser false, but got %v", anwser)
	}
}

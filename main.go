package main

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for right >= left {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return target
		} else if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func main() {
}

package main

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for right >= left {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// TODO: 搞懂以下兩者差異
func SumTwoInt(num1 int, num2 int) int {
	anwser := num1 + num2
	return anwser
}

/*
func SumTwoInt(num1 int, num2 int) int {
    return num1 + num2
}
*/

func TwoSum(nums []int, target int) []int {
	anwser := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				anwser = append(anwser, i, j)
			}
		}
	}
	return anwser
}

func main() {
}

package main

import "fmt"

//二分查找
func BinarySearch(key int, nums [9]int) int {
	var low int = 0
	var high int = len(nums) - 1
	var mid int = low/2 + high/2

	for ; low <= high; mid = low + (high-low)/2 {
		if key == nums[mid] {
			return mid
		}
		if key < nums[mid] {
			high = mid - 1
			continue
		}
		if key > nums[mid] {
			low = mid + 1
			continue
		}
	}

	return -1
}

func main() {
	n := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	index := BinarySearch(10, n)
	fmt.Println(index)
}

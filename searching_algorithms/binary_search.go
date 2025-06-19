package main

import "fmt"

// binarySearch searches for a target value in a sorted slice.
// It returns the index of the target if found, otherwise -1.
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid // Target found
		}

		if nums[mid] > target {
			high = mid - 1 // Search in the left half
		} else {
			low = mid + 1 // Search in the right half
		}
	}

	return -1 // Target not found
}

func main() {
	nums := []int{0, 1, 2, 3, 4}
	target := 3

	index := binarySearch(nums, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}

package main

import (
	"fmt"
)

func selectionSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			nums[minIdx], nums[i] = nums[i], nums[minIdx]
		}
	}
}

func main() {
	nums := []int{5, 6, 3, 4, 2, 1}
	selectionSort(nums)
	fmt.Println(nums)
}
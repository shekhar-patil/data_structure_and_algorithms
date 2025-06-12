package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 6, 3, 8, 2, 1}

	for i := 0; i < len(nums); i++ {
		minIdx := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		nums[minIdx], nums[i] = nums[i], nums[minIdx]
	}

	fmt.Println(nums)
}
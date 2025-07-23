package main

import (
  "fmt"
)

func binarySearch(nums []int, target int) int {
  low, high := 0, len(nums) - 1

    for low <= high {
        mid := (low + high) / 2

        if nums[mid] == target {
        return mid
        }

        if nums[mid] > target {
        high = mid - 1
        } else {
            low = mid + 1
        }
    }

  return -1
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
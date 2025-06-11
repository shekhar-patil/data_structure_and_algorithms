package main

import (
  "fmt"
)

func insertionSort(nums []int) {
  for i := 1; i < len(nums); i++ {
    key := nums[i]
    j := i - 1

    for j >= 0 && nums[j] > key {
      nums[j+1] = nums[j]
      j--
    }
    nums[j+1] = key
  }
}

func main() {
  nums := []int{10, 2, 5, 3, 1}
  insertionSort(nums)
  fmt.Println(nums)
}
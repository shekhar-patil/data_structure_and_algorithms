package main

import (
  "fmt"
)

func main() {
  nums := []int{5, 6, 3, 2, 4, 1}

  for i := 1; i < len(nums); i++ {
    key := nums[i]
    j := i - 1

    for j >= 0 && nums[j] > key {
      nums[j+1] = nums[j]
      j--
    }

    nums[j+1] = key
  }

  fmt.Println(nums)
}

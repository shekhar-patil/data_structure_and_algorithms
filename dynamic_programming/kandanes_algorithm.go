package main

import (
  "fmt"
  "math"
)

func maxSubArray(nums []int) int {
  maxSum := math.MinInt64
  sum := 0

  for i := 1; i < len(nums); i++ {
    sum += nums[i]

    if sum > maxSum {
      maxSum = sum
    }

    if sum < 0 {
      sum = 0
    }
  }

  return maxSum
}

func main() {
  arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
  fmt.Println("Maximum Subarray Sum:", maxSubArray(arr))
}

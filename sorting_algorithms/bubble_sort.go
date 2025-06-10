package main

import (
 "fmt"
)

func BubbleSort(nums []int) {
    n := len(nums)
    
    for i := 0; i < n; i++ {
        swapped := false
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] {
                swapped = true
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }

        if !swapped {
            break
        }
    }

    fmt.Println(nums)
}

func main() {
    nums := []int{10, 2, 5, 3, 1}
    BubbleSort(nums)
}
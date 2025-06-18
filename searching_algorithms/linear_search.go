package main

import "fmt"

func linearSearch(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i
        }
    }
    return -1 // Not found
}

func main() {
    nums := []int{5, 2, 9, 4, 7}
    target := 4

    index := linearSearch(nums, target)
    if index != -1 {
        fmt.Printf("Element found at index %d\n", index)
    } else {
        fmt.Println("Element not found")
    }
}
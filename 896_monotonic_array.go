package main

import (
	"fmt"
)

func TestMonotonicArray() {
	fmt.Println(isMonotonic([]int{}) == true)
	fmt.Println(isMonotonic([]int{1}) == true)
	fmt.Println(isMonotonic([]int{2, 1}) == true)
	fmt.Println(isMonotonic([]int{1, 1, 1}) == true)
	fmt.Println(isMonotonic([]int{1, 1, 2}) == true)
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}) == true)
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}) == true)
	fmt.Println(isMonotonic([]int{1, 2, 1}) == false)
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	// confirm monotone increasing or monotone decreasing
	flag := false // increasing
	i := 1
	for ; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			flag = false // increasing
			break
		} else if nums[i-1] > nums[i] {
			flag = true // decreasing
			break
		}
	}
	// all numbers are equal
	if i == len(nums) {
		return true
	}

	for ; i < len(nums); i++ {
		if !flag { // check increasing
			if nums[i-1] > nums[i] {
				return false
			}
		}

		if flag { // check decreasing
			if nums[i-1] < nums[i] {
				return false
			}
		}
	}
	return true
}

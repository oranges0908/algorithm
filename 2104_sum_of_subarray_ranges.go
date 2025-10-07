package main

import (
	"container/list"
	"fmt"
)

func TestSumOfSubArrayRanges() {
	fmt.Println(sumOfSubarrayRanges([]int{}) == 0)
	fmt.Println(sumOfSubarrayRanges([]int{1}) == 0)
	fmt.Println(sumOfSubarrayRanges([]int{1, 2, 3}) == 4)
	fmt.Println(sumOfSubarrayRanges([]int{1, 3, 3}) == 4)
	fmt.Println(sumOfSubarrayRanges([]int{4, -2, -3, 4, 1}) == 59)
}

func sumOfSubarrayRanges(nums []int) int {
	//sum of min
	minLeftArr := make([]int, len(nums))
	l := list.New()
	for i := 0; i < len(nums); i++ {
		for l.Len() > 0 && l.Back().Value.(element2).value >= nums[i] {
			l.Remove(l.Back())
		}
		if l.Len() > 0 {
			minLeftArr[i] = l.Back().Value.(element2).index
		} else {
			minLeftArr[i] = -1
		}
		l.PushBack(element2{index: i, value: nums[i]})
	}

	minRightArr := make([]int, len(nums))
	l = list.New()
	for i := len(nums) - 1; i >= 0; i-- {
		for l.Len() > 0 && l.Back().Value.(element2).value > nums[i] {
			l.Remove(l.Back())
		}
		if l.Len() > 0 {
			minRightArr[i] = l.Back().Value.(element2).index
		} else {
			minRightArr[i] = len(nums)
		}
		l.PushBack(element2{index: i, value: nums[i]})
	}

	minimum := 0
	for i := 0; i < len(nums); i++ {
		minimum += (i - minLeftArr[i]) * (minRightArr[i] - i) * nums[i]
	}

	//sum of max
	maxLeftArr := make([]int, len(nums))
	l = list.New()
	for i := 0; i < len(nums); i++ {
		for l.Len() > 0 && l.Back().Value.(element2).value <= nums[i] {
			l.Remove(l.Back())
		}
		if l.Len() > 0 {
			maxLeftArr[i] = l.Back().Value.(element2).index
		} else {
			maxLeftArr[i] = -1
		}
		l.PushBack(element2{index: i, value: nums[i]})
	}

	maxRightArr := make([]int, len(nums))
	l = list.New()
	for i := len(nums) - 1; i >= 0; i-- {
		for l.Len() > 0 && l.Back().Value.(element2).value < nums[i] {
			l.Remove(l.Back())
		}
		if l.Len() > 0 {
			maxRightArr[i] = l.Back().Value.(element2).index
		} else {
			maxRightArr[i] = len(nums)
		}
		l.PushBack(element2{index: i, value: nums[i]})
	}

	maximum := 0
	for i := 0; i < len(nums); i++ {
		maximum += (i - maxLeftArr[i]) * (maxRightArr[i] - i) * nums[i]
	}

	// difference
	return maximum - minimum
}

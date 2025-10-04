package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func TestNextGreaterElement() {
	fmt.Println(reflect.DeepEqual(nextGreaterElement1([]int{4, 1, 2}, []int{1, 3, 4, 2}), []int{-1, 3, -1}))
	fmt.Println(reflect.DeepEqual(nextGreaterElement1([]int{2, 4}, []int{1, 2, 3, 4}), []int{3, -1}))
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int, len(nums2))

	mm := list.List{}
	for i := len(nums2) - 1; i >= 0; i-- {
		// pop the lease values
		for mm.Len() > 0 && mm.Back().Value.(int) < nums2[i] {
			mm.Remove(mm.Back())
		}
		// calculate return value
		if mm.Len() > 0 {
			nextGreater[nums2[i]] = mm.Back().Value.(int)
		} else {
			nextGreater[nums2[i]] = -1
		}
		mm.PushBack(nums2[i])
	}

	rc := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if v, ok := nextGreater[nums1[i]]; ok {
			rc[i] = v
		} else {
			rc[i] = -1
		}
	}
	return rc
}

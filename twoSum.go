package main

import "fmt"

func testTwoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))  //[0,1]
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 10)) //[]
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 7))  //[]
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, num := range nums {
		cache[num] = i
	}

	for i, num := range nums {
		if v, ok := cache[target-num]; ok {
			if v != i {
				return []int{i, v}
			}
		}
	}
	return []int{}
}

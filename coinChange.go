package main

import "fmt"

func testCoinChange() {
	fmt.Println(coinChange([]int{}, 0) == 0)
	fmt.Println(coinChange([]int{}, 1) == -1)
	fmt.Println(coinChange([]int{1}, 0) == 0)
	fmt.Println(coinChange([]int{1}, 1) == 1)
	fmt.Println(coinChange([]int{1}, 2) == 2)
	fmt.Println(coinChange([]int{1, 2}, 2) == 1)
	fmt.Println(coinChange([]int{1, 2, 5}, 11) == 3)
}

func coinChange(coins []int, amount int) int {
	dpTable := make(map[int]int, amount+1)
	for i := 0; i <= amount; i++ {
		coinChangeCore(coins, i, dpTable)
	}
	if v, ok := dpTable[amount]; ok {
		return v
	}
	return -1
}

func coinChangeCore(coins []int, amount int, dpTable map[int]int) int {
	// find from cache
	if v, ok := dpTable[amount]; ok {
		return v
	}

	rc := -1
	defer func() {
		// cache
		dpTable[amount] = rc
	}()

	// boundary
	if amount == 0 {
		rc = 0
		return rc
	}
	if amount < 0 {
		rc = -1
		return rc
	}

	// traversal to find the minimum number of coins
	min := -1
	for _, coin := range coins {
		sb := coinChangeCore(coins, amount-coin, dpTable)
		if sb == -1 {
			continue
		}
		number := 1 + sb
		if number < min || min < 0 {
			min = number
		}
	}
	rc = min
	return rc
}

package main

import "fmt"

func testMinChanges() {
	fmt.Println(minChanges("1001") == 2)
	fmt.Println(minChanges("10") == 1)
	fmt.Println(minChanges("0000") == 0)
}

func minChanges(s string) int {
	if len(s)%2 != 0 {
		return -1
	}
	count := 0
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			count++
		}
	}
	return count
}

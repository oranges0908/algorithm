package main

import (
	"fmt"
)

func TestUniqueSubstringsInWraparoundString() {

	fmt.Println(uniqueSubstringsInWraparoundString2("a") == 1)
	fmt.Println(uniqueSubstringsInWraparoundString2("cac") == 2)
	fmt.Println(uniqueSubstringsInWraparoundString2("ab") == 3)
	fmt.Println(uniqueSubstringsInWraparoundString2("zab") == 6)
}

func uniqueSubstringsInWraparoundString(s string) int {
	cache := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		sub := s[i : i+1]
		cache[sub] = checkInBase(sub)
		for j := i + 2; j <= len(s); j++ {
			sub := s[i:j]
			if _, ok := cache[sub]; !ok {
				cache[sub] = checkInBase(sub[len(sub)-2:])
			}
			if cache[sub] {
				continue
			}
			break
		}
	}
	total := 0
	for _, v := range cache {
		if v {
			total++
		}
	}
	return total
}

func uniqueSubstringsInWraparoundString2(s string) int {

	return 0
}

func checkInBase(a string) bool {
	if len(a) == 0 {
		return false
	}
	if len(a) == 1 {
		return true
	}
	df := int(a[1]) - int(a[0])
	if df == 1 || df == -25 {
		return true
	}
	return false
}

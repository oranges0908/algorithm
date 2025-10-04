package main

import (
	"fmt"
)

func TestRepeatedSubstringPattern() {
	fmt.Println(repeatedSubstringPattern2("a") == false)
	fmt.Println(repeatedSubstringPattern2("aa") == true)
	fmt.Println(repeatedSubstringPattern2("aba") == false)
	fmt.Println(repeatedSubstringPattern2("abab") == true)
	fmt.Println(repeatedSubstringPattern2("abcabcabcabc") == true)
	fmt.Println(repeatedSubstringPattern2("abcababcab") == true)
	fmt.Println(repeatedSubstringPattern2("babbabbabbabbab") == true)
}

func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	}
	checking := false
	compareCursor := 0
	checkCursor := 0
	for i := 1; i < len(s); i++ {
		if !checking && s[i] == s[compareCursor] {
			checking = true
			checkCursor = i
		}
		if checking {
			if s[i] == s[compareCursor] {
				compareCursor++
			} else {
				checking = false
				compareCursor = 0
				i = checkCursor
			}
		}
	}
	if checking && len(s)%(len(s)-compareCursor) == 0 {
		return true
	}
	return false
}

func repeatedSubstringPattern2(s string) bool {
	if len(s) < 2 {
		return false
	}
	checkCursors := make([]int, 0)
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			checkCursors = append(checkCursors, i)
		}
	}

	for _, checkCursor := range checkCursors {
		i := 0
		j := checkCursor
		for j < len(s) {
			if s[i] != s[j] {
				break
			}
			i++
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false
}

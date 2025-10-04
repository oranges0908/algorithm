package main

import "fmt"

func testLengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("") == 0)
	fmt.Println(lengthOfLongestSubstring(" ") == 1)
}

func lengthOfLongestSubstring(s string) int {
	cache := make(map[string]int)
	maxLen := 0
	for j := 0; j < len(s); j++ {
		// repeated
		if cur, ok := cache[string(s[j])]; ok {
			// record max len
			if len(cache) > maxLen {
				maxLen = len(cache)
			}
			// pop the character before(include) repeated character
			for k, v := range cache {
				if v <= cur {
					delete(cache, k)
				}
			}
		}
		// no repeat, set new character
		cache[string(s[j])] = j
	}
	if len(cache) > maxLen {
		maxLen = len(cache)
	}
	return maxLen
}

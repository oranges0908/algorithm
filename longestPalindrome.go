package main

import "fmt"

func testLongestPalindrome() {
	fmt.Println(longestPalindrome("") == "")
	fmt.Println(longestPalindrome("a") == "a")
	fmt.Println(longestPalindrome("a1") == "a")
	fmt.Println(longestPalindrome("aa") == "aa")
	fmt.Println(longestPalindrome("aaa") == "aaa")
	fmt.Println(longestPalindrome("aaba") == "aba")
	fmt.Println(longestPalindrome("aabba") == "abba")
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	i := 0
	j := 0
	cur_i, cur_j := 0, 0
	maxlen := 0
	maxlen_i := 0
	maxlen_j := 0
	for {
		// bound check
		if i < len(s) && j < len(s) {
			// center check
			if s[i] == s[j] {
				cur_i = i
				cur_j = j
				if maxlen < cur_j-cur_i+1 {
					maxlen = cur_j - cur_i + 1
					maxlen_i = cur_i
					maxlen_j = cur_j
				}

				// extend both side,find longest Palindrome
				for {
					cur_i--
					cur_j++
					if cur_i >= 0 && cur_j < len(s) {
						if s[cur_i] == s[cur_j] {
							continue
						}
					}
					if maxlen < cur_j-cur_i-1 {
						maxlen = cur_j - cur_i - 1
						maxlen_i = cur_i + 1
						maxlen_j = cur_j - 1
					}
					break
				}
			}
			// move on
			if i == j {
				j++
			} else {
				i++
			}
		} else {
			break
		}
	}
	if maxlen_i < maxlen_j {
		return s[maxlen_i : maxlen_j+1]
	} else {
		return string(s[maxlen_i])
	}
}

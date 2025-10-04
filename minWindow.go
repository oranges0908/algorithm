package main

import "fmt"

func testMinWindow() {
	fmt.Println(minWindow("", "") == "")
	fmt.Println(minWindow("", "a") == "")
	fmt.Println(minWindow("a", "") == "")
	fmt.Println(minWindow("a", "a") == "a")
	fmt.Println(minWindow("aa", "a") == "a")
	fmt.Println(minWindow("a", "aa") == "")
	fmt.Println(minWindow("abc", "ac") == "abc")
	fmt.Println(minWindow("abc", "bc") == "bc")
	fmt.Println(minWindow("acbbaca", "aba") == "baca")
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd") == "abbbbbcdd")

}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	target := make(map[string]int)
	for i := 0; i < len(t); i++ {
		target[string(t[i])]++
	}

	locations := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := target[string(s[i])]; ok {
			locations = append(locations, i)
		}
	}

	if len(locations) < len(t) {
		return ""
	}

	minStr := ""
	start := 0
	end := len(t) - 1

	cache := make(map[string]int)
	for {
		// check result
		if len(cache) == 0 {
			for i := start; i <= end; i++ {
				cache[string(s[locations[i]])]++
			}
		}

		// check end
		succ := true
		if len(cache) == len(target) {
			for char, count := range target {
				if v, ok := cache[string(char)]; !ok {
					succ = false
					break
				} else {
					if v < count {
						succ = false
						break
					}
				}
			}
		} else {
			succ = false
		}

		if succ {
			str := s[locations[start]:]
			if locations[end]+1 < len(s) {
				str = s[locations[start] : locations[end]+1]
			}
			if len(str) < len(minStr) || minStr == "" {
				minStr = str
			}
			// move start
			if start < end {
				cache[string(s[locations[start]])]--
				if cache[string(s[locations[start]])] <= 0 {
					delete(cache, string(s[locations[start]]))
				}
				start++
			} else {
				if end < len(locations)-1 {
					end++
					cache[string(s[locations[end]])]++
				} else {
					break
				}
			}
		} else {
			// move end
			if end < len(locations)-1 {
				end++
				cache[string(s[locations[end]])]++
			} else {
				break
			}
		}
	}
	return minStr
}

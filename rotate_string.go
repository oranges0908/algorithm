package main

func testRotateString() {
	println(rotateString("abcde", "cdeab"))
	println(rotateString("", ""))
	println(rotateString("abcde", "cdeab1"))
	println(rotateString("abcde", "cdeaf"))
	println(rotateString("a", "a"))
	println(rotateString("ab", "ba"))
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 && len(goal) == 0 {
		return true
	}
	for i := 0; i < len(goal); i++ {
		// find header
		if goal[i] == s[0] {
			cur := 1
			for j := i + 1; ; j++ {
				if j >= len(goal) {
					j = 0
				}
				if cur == len(s) {
					return true // match
				}
				if goal[j] != s[cur] {
					break // not match
				}
				cur++
			}
		}
	}
	// not match
	return false
}

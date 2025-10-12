package main

import "fmt"

func TestMaxNestingDepthOf2ValidParenthesesStrings() {
	fmt.Println(maxNestingDepthOf2ValidParenthesesStrings("(()())"))
	//Output: [0,1,1,1,1,0]
	fmt.Println(maxNestingDepthOf2ValidParenthesesStrings("()(())()"))
	//Output: [0,0,0,1,1,0,1,1]
	fmt.Println(maxNestingDepthOf2ValidParenthesesStrings("(((()))((())))"))
}

func maxNestingDepthOf2ValidParenthesesStrings(s string) []int {
	d := 0
	rc := make([]int, len(s))
	for i, c := range s {
		if c == '(' {
			rc[i] = d % 2
			d++
		} else {
			d--
			rc[i] = d % 2
		}
	}

	return rc
}

func maxNestingDepthOf2ValidParenthesesStrings1(s string) []int {
	depth := 0

	d := 0
	rc := make([]int, len(s))
	for i, c := range s {
		if c == '(' {
			d++
			rc[i] = d
			if d > depth {
				depth = d
			}
		} else {
			rc[i] = d
			d--
		}
	}

	for i := 0; i < len(rc); i++ {
		rc[i] *= 2
		rc[i] /= (depth + 1)
	}
	return rc
}

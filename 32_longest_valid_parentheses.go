package main

import (
	"container/list"
	"fmt"
)

func TestLongestValidParentheses() {
	fmt.Println(longestValidParentheses("") == 0)
	fmt.Println(longestValidParentheses("(") == 0)
	fmt.Println(longestValidParentheses("()") == 2)
	fmt.Println(longestValidParentheses("(()") == 2)
	fmt.Println(longestValidParentheses(")()())") == 4)
	fmt.Println(longestValidParentheses("())()())") == 4)
}

func longestValidParentheses(s string) int {
	l := 0
	p := []int{}
	for i, c := range s {
		if c == ')' {
			if len(p) > 0 && s[p[len(p)-1]] == '(' {
				p = p[:len(p)-1]

				r := i + 1
				if len(p) > 0 {
					r = i - p[len(p)-1]
				}
				if r > l {
					l = r
				}
				continue
			}
		}
		p = append(p, i)
	}
	return l
}

func longestValidParentheses1(s string) int {
	minStack := list.New()
	parenthesesStack := list.New()

	for i := 0; i < len(s); i++ {
		if parenthesesStack.Len() > 0 && (parenthesesStack.Back().Value.(int32) == '(' && s[i] == ')') {
			parenthesesStack.Remove(parenthesesStack.Back())

			for minStack.Len() > 0 && minStack.Back().Value.(element2).value >= parenthesesStack.Len() {
				minStack.Remove(minStack.Back())
			}
			minStack.PushBack(element2{value: parenthesesStack.Len(), index: i})
			continue
		}
		parenthesesStack.PushBack(int32(s[i]))
	}

	rc := 0
	sumrc := 0
	for minStack.Len() > 0 {
		v := minStack.Front().Value.(element2)
		minStack.Remove(minStack.Front())

		r := v.index + 1 - v.value - sumrc
		if r > rc {
			rc = r
		}
		sumrc += r
	}
	return rc
}

package main

import (
	"container/list"
	"fmt"
)

func TestValidParentheses() {
	fmt.Println(validParentheses("") == true)
	fmt.Println(validParentheses("()") == true)
	fmt.Println(validParentheses("()[]{}") == true)
	fmt.Println(validParentheses("(]") == false)
	fmt.Println(validParentheses("([])") == true)
	fmt.Println(validParentheses("([)]") == false)
}

func validParentheses(s string) bool {
	l := list.New()

	for _, c := range s {
		if l.Len() > 0 {
			v := l.Back().Value.(int32)
			if checkParenthesesCompare(v, c) {
				l.Remove(l.Back())
			} else {
				l.PushBack(c)
			}
		} else {
			l.PushBack(c)
		}
	}

	return l.Len() == 0
}

func checkParenthesesCompare(a, b int32) bool {
	if a == '(' && b == ')' || a == '{' && b == '}' || a == '[' && b == ']' {
		return true
	}
	return false
}

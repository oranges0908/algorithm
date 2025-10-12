package main

import (
	"container/list"
	"fmt"
)

func TestRemoveInvalidParenthese() {
	fmt.Println(removeInvalidParentheses(")"))
	fmt.Println(removeInvalidParentheses(")("))
	fmt.Println(removeInvalidParentheses("())"))
	fmt.Println(removeInvalidParentheses("())("))
	fmt.Println(removeInvalidParentheses("()())"))
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
}

func removeInvalidParentheses(s string) []string {
	// === 1️⃣ Use a stack to find all unmatched parentheses ===
	l := list.New()
	for i := 0; i < len(s); i++ {
		// Match and remove pairs like "()", otherwise keep track
		if l.Len() > 0 && l.Back().Value.(element2).value == '(' && s[i] == ')' {
			l.Remove(l.Back())
			continue
		}
		if s[i] == '(' || s[i] == ')' {
			l.PushBack(element2{value: int(s[i]), index: i})
		}
	}

	// === 2️⃣ Initialize a set of strings to process ===
	m1 := make(map[string]struct{})
	m1[s] = struct{}{} // start from the original string

	shift := 0 // track how many characters have been removed

	// === 3️⃣ Iteratively remove extra parentheses ===
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(element2)
		var si, ei int

		// Define removal range
		if v.value == int(')') {
			// For extra ')': remove any ')' before or at its position
			si = 0
			ei = v.index - shift
		} else {
			// For extra '(': remove any '(' after or at its position
			si = v.index - shift
			ei = len(s) - shift - 1
		}
		shift++

		m2 := make(map[string]struct{})

		// For each string in current set, remove one invalid parenthesis
		for str := range m1 {
			for i := si; i <= ei; i++ {
				if int(str[i]) == v.value {
					// Remove the parenthesis at position i
					newStr := str[:i] + str[i+1:]
					m2[newStr] = struct{}{}
				}
			}
		}
		m1 = m2 // proceed to next iteration
	}

	rc := make([]string, 0)
	for k, _ := range m1 {
		rc = append(rc, k)
	}
	return rc
}

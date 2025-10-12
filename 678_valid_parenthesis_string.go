package main

import (
	"container/list"
	"fmt"
)

func TestValidParenthesisString() {
	fmt.Println(validParenthesisString("()") == true)
	fmt.Println(validParenthesisString("*(") == false)
	fmt.Println(validParenthesisString("*)") == true)
	fmt.Println(validParenthesisString("(*)") == true)
	fmt.Println(validParenthesisString("(*))") == true)
	fmt.Println(validParenthesisString("(((((*)(*)*)))))))((**)))))") == false)
	fmt.Println(validParenthesisString("(((((*(((((*((**(((*)*((((**))*)*)))))))))((*(((((**(**)") == false)
	fmt.Println(validParenthesisString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())") == false)
	fmt.Println(validParenthesisString("((((((((*)*)))))))))((") == false)
	fmt.Println(validParenthesisString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()") == true)
}

// 区间+贪心 O(n)
func validParenthesisString(s string) bool {
	low, high := 0, 0
	for _, c := range s {
		if c == '(' {
			low++
			high++
		} else if c == ')' {
			if low > 0 {
				low--
			}
			high--
		} else { // '*'
			if low > 0 {
				low--
			}
			high++
		}
		if high < 0 { // too many ')'
			return false
		}
	}
	return low == 0
}

// 双栈法 O(n)
func validParenthesisString3(s string) bool {
	left, star := []int{}, []int{}
	// Pair ')' with early '(' or '*'
	for i, c := range s {
		if c == '(' {
			left = append(left, i)
		} else if c == '*' {
			star = append(star, i)
		} else { // ')'
			if len(left) > 0 {
				left = left[:len(left)-1]
			} else if len(star) > 0 {
				star = star[:len(star)-1]
			} else {
				return false
			}
		}
	}

	// Pair remaining '(' with later '*'
	for len(left) > 0 && len(star) > 0 {
		if left[len(left)-1] > star[len(star)-1] {
			return false
		}
		left = left[:len(left)-1]
		star = star[:len(star)-1]
	}

	return len(left) == 0
}

// stack + queue, time complexity O(n^2)
func validParenthesisString2(s string) bool {
	l := list.New()
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			// ')' is more than '(' and '*'，string is invalid
			v := l.Back()
			if v == nil {
				return false
			}

			// find '('
			for v.Value.(int) == int('*') {
				v = v.Prev()
				if v == nil {
					break
				}
			}

			if v != nil {
				l.Remove(v) // remove '(' prioritize
			} else {
				l.Remove(l.Back()) // remove '*' alternative
			}
			continue
		}
		l.PushBack(int(s[i])) // push '(' '*' into stack
	}

	lb := 0
	for l.Len() > 0 {
		if l.Front().Value.(int) == '(' { // accumulate '('
			lb++
		} else if l.Front().Value.(int) == '*' {
			if lb > 0 { // only take advantage of '*' which is after '('
				lb-- // Pair Elimination '(' and '*'
			}
		}
		l.Remove(l.Front())
	}
	if lb == 0 { // all the '(' and ')' are removed successfully, no '(' left
		return true
	}
	return false
}

// time complexity O(n*3^k)
func validParenthesisString1(s string) bool {

	l := list.New()
	l.PushBack(s)

loop:
	for l.Len() > 0 {
		v := l.Front().Value.(string)
		l.Remove(l.Front())
		cnt := 0
		for i := 0; i < len(v); i++ {
			if v[i] == '(' {
				cnt++
			} else if v[i] == ')' {
				cnt--
				if cnt < 0 {
					continue loop
				}
			} else {
				l.PushBack(v[:i] + v[i+1:])
				l.PushBack(v[:i] + "(" + v[i+1:])
				l.PushBack(v[:i] + ")" + v[i+1:])
				continue loop
			}
		}
		if cnt == 0 {
			return true
		}
	}
	return false
}

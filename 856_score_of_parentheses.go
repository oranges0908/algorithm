package main

import (
	"container/list"
	"fmt"
	"math"
)

func TestScoreOfParentheses() {
	fmt.Println(scoreOfParentheses("") == 0)
	fmt.Println(scoreOfParentheses("()") == 1)
	fmt.Println(scoreOfParentheses("()()") == 2)
	fmt.Println(scoreOfParentheses("()()()") == 3)
	fmt.Println(scoreOfParentheses("(())") == 2)
	fmt.Println(scoreOfParentheses("((()))") == 4)
	fmt.Println(scoreOfParentheses("(()())") == 4)
	fmt.Println(scoreOfParentheses("(()()())") == 6)
	fmt.Println(scoreOfParentheses("(()())(())") == 6)
	fmt.Println(scoreOfParentheses("(()())(())()") == 7)
}

func scoreOfParentheses(s string) int {
	stk := []int{}

	for _, c := range s {
		if c == '(' {
			stk = append(stk, math.MinInt)
			continue
		}

		if stk[len(stk)-1] == math.MinInt {
			stk[len(stk)-1] = 1
		} else {
			tm := 0
			for len(stk) > 0 && stk[len(stk)-1] != math.MinInt {
				tm += stk[len(stk)-1]
				stk = stk[:len(stk)-1]
			}
			if len(stk) > 0 {
				stk[len(stk)-1] = tm * 2
			} else {
				stk = append(stk, tm*2)
			}
		}
	}

	rc := 0
	for len(stk) > 0 {
		rc += stk[len(stk)-1]
		stk = stk[:len(stk)-1]
	}
	return rc
}

type element3 struct {
	start int
	end   int
	value int
}

func scoreOfParentheses1(s string) int {
	stk := []int{}

	l := list.New()
	for i, c := range s {
		if c == '(' {
			stk = append(stk, i)
			continue
		}

		p := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		nv := element3{value: 1, start: p, end: i}
		// 计算相乘
		if p < i-1 {
			last := l.Back().Value.(element3)
			nv.value = 2 * last.value
			l.Remove(l.Back())
		}

		// 计算前向的相加
		if l.Len() > 0 {
			last := l.Back().Value.(element3)
			if last.end == p-1 {
				nv.value += last.value
				nv.start = last.start
				l.Remove(l.Back())
			}
		}
		l.PushBack(nv)

	}

	if l.Len() > 0 {
		return l.Front().Value.(element3).value
	}
	return 0
}

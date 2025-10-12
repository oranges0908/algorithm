package main

import "fmt"

func TestInsertions2BalanceParenthesesString() {
	fmt.Println(minInsertions2BalanceParenthesesString("") == 0)
	fmt.Println(minInsertions2BalanceParenthesesString("(") == 2)
	fmt.Println(minInsertions2BalanceParenthesesString(")") == 2)
	fmt.Println(minInsertions2BalanceParenthesesString("))") == 1)
	fmt.Println(minInsertions2BalanceParenthesesString("(()))") == 1)
	fmt.Println(minInsertions2BalanceParenthesesString("())") == 0)
	fmt.Println(minInsertions2BalanceParenthesesString("))())(") == 3)
	fmt.Println(minInsertions2BalanceParenthesesString("()()()()()(") == 7)

}

func minInsertions2BalanceParenthesesString(s string) int {
	need := 0
	insertCount := 0
	for _, c := range s {
		if c == '(' {
			need += 2

			if need%2 == 1 {
				need--
				insertCount++
			}

			continue
		}
		if c == ')' {
			need--
			if need < 0 {
				insertCount++
				need += 2
			}
		}
	}
	return insertCount + need
}

// 双栈，线性思维，缺了立刻补充
func minInsertions2BalanceParenthesesString1(s string) int {
	insertCount := 0

	lb := []int{}
	rb := []int{}
	doubleCheck := false
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if doubleCheck {
				insertCount++
				doubleCheck = false
			}
			lb = append(lb, i)
			continue
		}
		if s[i] == ')' {
			if doubleCheck {
				doubleCheck = false
			} else {
				rb = append(rb, i)
				doubleCheck = true
			}
			continue
		}
	}
	if doubleCheck {
		insertCount++
	}

	for len(lb) > 0 && len(rb) > 0 {
		if rb[len(rb)-1] > lb[len(lb)-1] {
			rb = rb[:len(rb)-1]
			lb = lb[:len(lb)-1]
			continue
		} else {
			lb = lb[:len(lb)-1]
			insertCount += 2
		}
	}

	if len(rb) > 0 {
		insertCount += len(rb)
	}
	if len(lb) > 0 {
		insertCount += len(lb) * 2
	}
	return insertCount
}

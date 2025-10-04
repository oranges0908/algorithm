package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func testAddTwoNumbers() {
	testAddTwoNumbers1()
	testAddTwoNumbers2()
	testAddTwoNumbers3()
}

func testAddTwoNumbers1() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	ep := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}
	rc := addTwoNumbers(l1, l2)
	checkRc(ep, rc)
}

func testAddTwoNumbers2() {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}
	ep := &ListNode{Val: 0}
	rc := addTwoNumbers(l1, l2)
	checkRc(ep, rc)
}

func testAddTwoNumbers3() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}
	ep := &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}}
	rc := addTwoNumbers(l1, l2)
	checkRc(ep, rc)
}

func checkRc(l1 *ListNode, l2 *ListNode) {
	i := l1
	j := l2
	for i != nil && j != nil {
		if i.Val != j.Val {
			break
		}
		i = i.Next
		j = j.Next
	}
	if i != nil || j != nil {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1_cur := l1
	l2_cur := l2
	rc_cur := &ListNode{}
	rc_head := rc_cur
	up := 0
	for {
		if l1_cur != nil && l2_cur != nil {
			rc_cur.Next = &ListNode{}
			rc_cur = rc_cur.Next
			rc_cur.Val = l1_cur.Val + l2_cur.Val + up
			if rc_cur.Val >= 10 {
				up = 1
				rc_cur.Val -= 10
			} else {
				up = 0
			}
		} else {
			break
		}
		l1_cur = l1_cur.Next
		l2_cur = l2_cur.Next
	}

	if l1_cur != nil {
		for {
			rc_cur.Next = &ListNode{}
			rc_cur = rc_cur.Next
			rc_cur.Val = l1_cur.Val + up
			if rc_cur.Val >= 10 {
				up = 1
				rc_cur.Val -= 10
			} else {
				up = 0
			}
			l1_cur = l1_cur.Next
			if l1_cur == nil {
				break
			}
		}
	}

	if l2_cur != nil {
		for {
			rc_cur.Next = &ListNode{}
			rc_cur = rc_cur.Next
			rc_cur.Val = l2_cur.Val + up
			if rc_cur.Val >= 10 {
				up = 1
				rc_cur.Val -= 10
			} else {
				up = 0
			}
			l2_cur = l2_cur.Next
			if l2_cur == nil {
				break
			}
		}
	}

	if up > 0 {
		rc_cur.Next = &ListNode{}
		rc_cur = rc_cur.Next
		rc_cur.Val = up
		up = 0
	}

	return rc_head.Next
}

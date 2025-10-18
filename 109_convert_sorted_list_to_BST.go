package main

import "fmt"

func TestConvertSortedListToBST() {
	n := convertSortedListToBST(nil)
	fmt.Println("0", n == nil)

	n = convertSortedListToBST(&ListNode{Val: 1})
	fmt.Println("3", n.Val == 1)

	n = convertSortedListToBST(&ListNode{Val: 1, Next: &ListNode{Val: 3}})
	fmt.Println("2", n.Val == 1)
	fmt.Println("2", n.Right.Val == 3)

	n = convertSortedListToBST(&ListNode{Val: -10,
		Next: &ListNode{Val: -3,
			Next: &ListNode{Val: 0,
				Next: &ListNode{Val: 5,
					Next: &ListNode{Val: 9}}}}})
	fmt.Println("1", n.Val == 0)
	fmt.Println("1", n.Left.Val == -10)
	fmt.Println("1", n.Left.Right.Val == -3)
	fmt.Println("1", n.Right.Val == 5)
	fmt.Println("1", n.Right.Right.Val == 9)
}

func convertSortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	curListNode = head

	c := 0
	for curListNode != nil {
		c++
		curListNode = curListNode.Next
	}

	curListNode = head

	return sortedListToBST(0, c-1)
}

var curListNode *ListNode

func sortedListToBST(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2

	//中序遍历
	left := sortedListToBST(start, mid-1)

	rc := &TreeNode{Val: curListNode.Val}
	curListNode = curListNode.Next

	right := sortedListToBST(mid+1, end)

	rc.Left = left
	rc.Right = right

	return rc
}

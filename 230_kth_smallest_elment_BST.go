package main

import (
	"container/list"
	"fmt"
)

func TestFindKthSmallestBST() {
	n00 := &TreeNode{Val: 0}
	fmt.Println(findKthSmallestBST(n00, 1) == 0)

	n11 := &TreeNode{Val: 10}
	n00.Right = n11
	fmt.Println(findKthSmallestBST(n00, 1) == 0)
	fmt.Println(findKthSmallestBST(n00, 2) == 10)

	n00 = &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}
	fmt.Println(findKthSmallestBST(n00, 1) == 3)
	fmt.Println(findKthSmallestBST(n00, 2) == 7)
	fmt.Println(findKthSmallestBST(n00, 3) == 9)
	fmt.Println(findKthSmallestBST(n00, 4) == 15)
	fmt.Println(findKthSmallestBST(n00, 5) == 20)
}

func findKthSmallestBST(root *TreeNode, k int) int {
	cur := root
	l := list.New()
	for {
		if cur != nil {
			l.PushBack(cur)
			cur = cur.Left
			continue
		}

		if l.Back() == nil {
			return -1
		}
		cur = l.Back().Value.(*TreeNode)
		l.Remove(l.Back())
		k--
		if k == 0 {
			return cur.Val
		}

		cur = cur.Right
	}
}

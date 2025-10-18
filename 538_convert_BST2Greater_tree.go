package main

import (
	"container/list"
	"fmt"
)

func TestConvertBST2GreaterTree() {
	n00 := &TreeNode{Val: 5}
	fmt.Println(convertBST2GreaterTree(n00).Val == 5)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10}}
	c := convertBST2GreaterTree(n00)
	fmt.Println(c.Val == 15)
	fmt.Println(c.Right.Val == 10)

	n00 = &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 10}}
	fmt.Println(convertBST2GreaterTree(n00).Left.Val == 18)

	n00 = &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}
	fmt.Println(convertBST2GreaterTree(n00).Val == 51)

	n00 = &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 6}}
	fmt.Println(convertBST2GreaterTree(n00).Left.Val == 15)

	n00 = &TreeNode{Val: 20,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 22}, Right: &TreeNode{Val: 99}}}
	fmt.Println(convertBST2GreaterTree(n00).Right.Val == 147)
}

func convertBST2GreaterTree(root *TreeNode) *TreeNode {
	cc := 0
	// right - root - left
	l := list.New()
	cur := root
	for {
		if cur != nil {
			l.PushBack(cur)
			cur = cur.Right
			continue
		}

		if l.Len() == 0 {
			break
		}

		cur = l.Back().Value.(*TreeNode)
		l.Remove(l.Back())
		cur.Val += cc
		cc = cur.Val

		cur = cur.Left
	}
	return root
}

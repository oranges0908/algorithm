package main

import (
	"container/list"
	"fmt"
)

func TestIncreasingBST2orderSearchTree() {
	n00 := &TreeNode{Val: 5}
	fmt.Println(increasingBST2orderSearchTree(n00).Val == 5)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10}}
	fmt.Println(increasingBST2orderSearchTree(n00).Right.Val == 10)

	n00 = &TreeNode{Val: 5, Left: &TreeNode{Val: 2}}
	c := increasingBST2orderSearchTree(n00)
	fmt.Println(c.Val == 2)
	fmt.Println(c.Right.Val == 5)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 6}}}
	c = increasingBST2orderSearchTree(n00)
	fmt.Println(c.Val == 5)
	fmt.Println(c.Right.Val == 6)
	fmt.Println(c.Right.Right.Val == 10)

}

func increasingBST2orderSearchTree(root *TreeNode) *TreeNode {
	l := list.New()
	rc := &TreeNode{}
	mk := rc

	cur := root
	for {
		if cur != nil {
			l.PushBack(cur)
			cur = cur.Left
			continue
		}
		if l.Len() == 0 {
			break
		}
		cur = l.Back().Value.(*TreeNode)
		mk.Right = &TreeNode{Val: cur.Val}
		mk = mk.Right
		l.Remove(l.Back())

		cur = cur.Right
	}
	return rc.Right
}

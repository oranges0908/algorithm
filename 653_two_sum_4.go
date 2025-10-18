package main

import (
	"container/list"
	"fmt"
)

func TestTwoSum4() {
	n00 := &TreeNode{Val: 5}
	fmt.Println(TwoSum4(n00, 1) == false)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10}}
	fmt.Println(TwoSum4(n00, 15) == true)

	n00 = &TreeNode{Val: 5, Left: &TreeNode{Val: 2}}
	fmt.Println(TwoSum4(n00, 7) == true)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 6}}}
	fmt.Println(TwoSum4(n00, 11) == true)
	fmt.Println(TwoSum4(n00, 21) == false)
	fmt.Println(TwoSum4(n00, 16) == true)
}

func TwoSum4(root *TreeNode, k int) bool {
	m := make(map[int]bool)

	var fd func(node *TreeNode) bool
	fd = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if m[k-node.Val] {
			return true
		}
		m[node.Val] = true
		return fd(node.Left) || fd(node.Right)
	}
	return fd(root)
}

func TwoSum4_(root *TreeNode, k int) bool {
	l := list.New()
	m := make(map[int]struct{})

	cur := root
	for {
		if cur != nil {
			l.PushBack(cur)
			cur = cur.Left
			continue
		}
		if l.Len() == 0 {
			return false
		}

		cur = l.Back().Value.(*TreeNode)
		l.Remove(l.Back())

		if _, ok := m[k-cur.Val]; ok {
			return true
		}

		m[cur.Val] = struct{}{}

		cur = cur.Right
	}
}

package main

import "fmt"

func testIsBalancedBST() {
	fmt.Println(isBalancedBST(nil) == true)

	n00 := &TreeNode{Val: 0}
	fmt.Println(isBalancedBST(n00) == true)

	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	fmt.Println(isBalancedBST(n00) == true)

	n21 := &TreeNode{Val: 20}
	n10.Right = n21
	fmt.Println(isBalancedBST(n00) == false)

	n00 = &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(isBalancedBST(n00) == true)

	n00 = &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3}}}
	fmt.Println(isBalancedBST(n00) == false)
}

func isBalancedBST(root *TreeNode) bool {
	_, ok := checkBalancedBST(root)
	return ok
}

func checkBalancedBST(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	if root.Left == nil && root.Right == nil {
		return 1, true
	}
	ok := true
	lt, rt := 0, 0
	if root.Left != nil {
		lt, ok = checkBalancedBST(root.Left)
		if !ok {
			return -1, false
		}
	}
	if root.Right != nil {
		rt, ok = checkBalancedBST(root.Right)
		if !ok {
			return -1, false
		}
	}

	if lt > rt {
		if lt-rt > 1 {
			return -1, false
		} else {
			return lt + 1, true
		}
	}
	if rt-lt > 1 {
		return -1, false
	} else {
		return rt + 1, true
	}
}

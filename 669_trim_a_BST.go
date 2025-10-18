package main

import "fmt"

func TestTrimBST() {
	n00 := &TreeNode{Val: 5}
	fmt.Println(trimBST(n00, 6, 9) == nil)
	fmt.Println(trimBST(n00, 1, 9).Val == 5)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10}}
	fmt.Println(trimBST(n00, 10, 10).Val == 10)

	n00 = &TreeNode{Val: 5, Left: &TreeNode{Val: 2}}
	fmt.Println(trimBST(n00, 1, 10).Val == 5)
	fmt.Println(trimBST(n00, 1, 3).Val == 2)

	n00 = &TreeNode{Val: 5, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 6}}}
	fmt.Println(trimBST(n00, 1, 8).Right.Val == 6)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	//root
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	//left
	if root.Left != nil {
		root.Left = trimBST(root.Left, low, high)
	}

	//right
	if root.Right != nil {
		root.Right = trimBST(root.Right, low, high)
	}

	return root
}

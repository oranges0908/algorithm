package main

import "fmt"

func TestSymmetricTree() {
	fmt.Println(symmetricTree(nil) == true)

	n00 := &TreeNode{Val: 0}
	fmt.Println(symmetricTree(n00) == true)

	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	fmt.Println(symmetricTree(n00) == false)

	n11 := &TreeNode{Val: 11}
	n00.Right = n11
	fmt.Println(symmetricTree(n00) == false)

	n11.Val = 10
	fmt.Println(symmetricTree(n00) == true)

	n20 := &TreeNode{Val: 20}
	n10.Left = n20
	fmt.Println(symmetricTree(n00) == false)

	n23 := &TreeNode{Val: 20}
	n11.Right = n23
	fmt.Println(symmetricTree(n00) == true)

	n11.Left = n23
	n11.Right = nil
	fmt.Println(symmetricTree(n00) == false)
}

func symmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetricTreeNodes(root.Left, root.Right)
}

func symmetricTreeNodes(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if !symmetricTreeNodes(left.Left, right.Right) {
		return false
	}

	if !symmetricTreeNodes(left.Right, right.Left) {
		return false
	}
	return true
}

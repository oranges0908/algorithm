package main

import "fmt"

func TestValidateBinarySearch() {
	fmt.Println(isValidBST(nil) == true)

	n00 := &TreeNode{Val: 0}
	fmt.Println(isValidBST(n00) == true)

	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	fmt.Println(isValidBST(n00) == false)
	n10.Val = 0
	fmt.Println(isValidBST(n00) == false)
	n10.Val = -2
	fmt.Println(isValidBST(n00) == true)

	n21 := &TreeNode{Val: 20}
	n10.Right = n21
	fmt.Println(isValidBST(n00) == false)
	n21.Val = -2
	fmt.Println(isValidBST(n00) == false)
	n21.Val = -1
	fmt.Println(isValidBST(n00) == true)

	n00 = &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(n00) == true)

	n00 = &TreeNode{Val: 120,
		Left: &TreeNode{Val: 70,
			Left:  &TreeNode{Val: 50, Left: &TreeNode{Val: 20}, Right: &TreeNode{Val: 55}},
			Right: &TreeNode{Val: 100, Left: &TreeNode{Val: 75}, Right: &TreeNode{Val: 110}}},
		Right: &TreeNode{Val: 140,
			Left:  &TreeNode{Val: 130, Left: &TreeNode{Val: 119}, Right: &TreeNode{Val: 135}},
			Right: &TreeNode{Val: 160, Left: &TreeNode{Val: 150}, Right: &TreeNode{Val: 200}}}}
	fmt.Println(isValidBST(n00) == false)
}

func isValidBST(root *TreeNode) bool {
	return validateBST(root, nil, nil)
}

func validateBinarySearch(root *TreeNode, leftParent, rightParent *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		if rightParent != nil && root.Left.Val <= rightParent.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if leftParent != nil && root.Right.Val >= leftParent.Val {
			return false
		}
	}
	return validateBinarySearch(root.Left, root, nil) && validateBinarySearch(root.Right, nil, root)
}

func validateBST(root *TreeNode, leftP, rightP *TreeNode) bool {
	if root == nil {
		return true
	}
	if leftP != nil && root.Val <= leftP.Val || rightP != nil && root.Val >= rightP.Val {
		return false
	}
	return validateBST(root.Left, leftP, root) && validateBST(root.Right, root, rightP)
}

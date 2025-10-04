package main

import "fmt"

func TestRecoverBTPreInorderOrder() {
	fmt.Println(CompareTree(nil, RecoverBTPreInorderOrder([]int{}, []int{})) == true)

	n00 := &TreeNode{Val: 0}
	fmt.Println(CompareTree(n00, RecoverBTPreInorderOrder([]int{0}, []int{0})) == true)

	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	fmt.Println(CompareTree(n00, RecoverBTPreInorderOrder([]int{0, 10}, []int{10, 0})) == true)

	n11 := &TreeNode{Val: 11}
	n00.Right = n11
	fmt.Println(CompareTree(n00, RecoverBTPreInorderOrder([]int{0, 10, 11}, []int{10, 0, 11})) == true)

	n20 := &TreeNode{Val: 20}
	n10.Left = n20
	fmt.Println(CompareTree(n00, RecoverBTPreInorderOrder([]int{0, 10, 20, 11}, []int{20, 10, 0, 11})) == true)

	n23 := &TreeNode{Val: 23}
	n11.Right = n23
	fmt.Println(CompareTree(n00, RecoverBTPreInorderOrder([]int{0, 10, 20, 11, 23}, []int{20, 10, 0, 11, 23})) == true)
}

func RecoverBTPreInorderOrder(preorder []int, inorder []int) *TreeNode {
	if len(inorder) != len(preorder) || len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{}
	for i, v := range inorder {
		if v == preorder[0] {
			node.Val = v
			if i > 0 && len(preorder[1:i+1]) > 1 {
				node.Left = RecoverBTPreInorderOrder(preorder[1:i+1], inorder[:i])
			}
			if i < len(inorder)-1 && len(preorder[i+1:]) > 1 {
				node.Right = RecoverBTPreInorderOrder(preorder[i+1:], inorder[i+1:])
			}
			break
		}
	}
	return node
}

package main

import "fmt"

func CompareTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		fmt.Println("a == nil || b == nil")
		return false
	}

	if a.Val != b.Val {
		fmt.Println("a.Val != b.Val")
		return false
	}

	if !CompareTree(a.Left, b.Left) {
		return false
	}

	if !CompareTree(a.Right, b.Right) {
		return false
	}
	return true
}

func TestRecoverBTInorderPostOrder() {
	fmt.Println(CompareTree(nil, RecoverBTInorderPostOrder([]int{}, []int{})) == true)

	n00 := &TreeNode{Val: 0}
	fmt.Println(CompareTree(n00, RecoverBTInorderPostOrder([]int{0}, []int{0})) == true)

	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	fmt.Println(CompareTree(n00, RecoverBTInorderPostOrder([]int{10, 0}, []int{10, 0})) == true)

	n11 := &TreeNode{Val: 11}
	n00.Right = n11
	fmt.Println(CompareTree(n00, RecoverBTInorderPostOrder([]int{10, 0, 11}, []int{10, 11, 0})) == true)

	n20 := &TreeNode{Val: 20}
	n10.Left = n20
	fmt.Println(CompareTree(n00, RecoverBTInorderPostOrder([]int{20, 10, 0, 11}, []int{20, 10, 11, 0})) == true)

	n23 := &TreeNode{Val: 23}
	n11.Right = n23
	fmt.Println(CompareTree(n00, RecoverBTInorderPostOrder([]int{20, 10, 0, 11, 23}, []int{20, 10, 23, 11, 0})) == true)
}

func RecoverBTInorderPostOrder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) || len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{}
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			node.Val = v
			if i > 0 {
				node.Left = RecoverBTInorderPostOrder(inorder[:i], postorder[:i])
			}
			if i < len(inorder)-1 {
				node.Right = RecoverBTInorderPostOrder(inorder[i+1:], postorder[i:len(postorder)-1])
			}
			break
		}
	}
	return node
}

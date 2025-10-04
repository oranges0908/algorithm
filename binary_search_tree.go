package main

import "fmt"

func testSearchBST() {
	fmt.Println(searchBST(nil, 0) == nil)

	n00 := &TreeNode{Val: 0}
	fmt.Println(searchBST(n00, 0) == n00)
	fmt.Println(searchBST(n00, 1) == nil)

	n10 := &TreeNode{Val: -1}
	n00.Left = n10
	fmt.Println(searchBST(n00, -1) == n10)
	n11 := &TreeNode{Val: 20}
	n00.Right = n11
	fmt.Println(searchBST(n00, 20) == n11)
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func testInsertBST() {
	fmt.Println(insertBST(nil, 0) != nil)
	fmt.Println(insertBST(nil, 0).Val == 0)

	fmt.Println(insertBST(&TreeNode{Val: 0}, -1).Left != nil)
	fmt.Println(insertBST(&TreeNode{Val: 0}, -1).Left.Val == -1)

	fmt.Println(insertBST(&TreeNode{Val: 0}, 1).Right != nil)
	fmt.Println(insertBST(&TreeNode{Val: 0}, 1).Right.Val == 1)
}

func insertBST(root *TreeNode, val int) *TreeNode {
	//root
	if root == nil {
		return &TreeNode{Val: val}
	}

	//left
	if root.Val > val {
		root.Left = insertBST(root.Left, val)
	}
	//right
	if root.Val < val {
		root.Right = insertBST(root.Right, val)
	}
	return root // return itself
}

func testDeleteBST() {
	fmt.Println(deleteBST(nil, 0) == nil)
	fmt.Println(deleteBST(&TreeNode{Val: 0}, 0) == nil)

	fmt.Println(deleteBST(&TreeNode{Val: 0, Left: &TreeNode{Val: -1}}, -1).Left == nil)
	fmt.Println(deleteBST(&TreeNode{Val: 0, Left: &TreeNode{Val: -2}}, -1).Left.Val == -2)
	fmt.Println(deleteBST(&TreeNode{Val: 0, Right: &TreeNode{Val: 1}}, 1).Right == nil)
	fmt.Println(deleteBST(&TreeNode{Val: 0, Left: &TreeNode{Val: -2, Right: &TreeNode{Val: -1}}}, -2).Left.Val == -1)

	fmt.Println(deleteBST(&TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}},
	}, 3).Left.Left.Val == 2)
	fmt.Println(deleteBST(&TreeNode{Val: 1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 8, Right: &TreeNode{Val: 9}}},
	}, 1).Right.Left.Val == 9)
}

func deleteBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := getBSTMinNode(root.Right)
		root.Val = minNode.Val
		root.Right = deleteBST(root.Right, minNode.Val)
	}
	if root.Val > val {
		root.Left = deleteBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = deleteBST(root.Right, val)
	}
	return root
}

func getBSTMinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return getBSTMinNode(root.Left)
}

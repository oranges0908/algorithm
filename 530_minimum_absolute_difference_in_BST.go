package main

import (
	"fmt"
	"math"
)

func TestMinimumAbsDifferenceBST() {
	n00 := &TreeNode{Val: 5}
	n00.Right = &TreeNode{Val: 10}
	fmt.Println(minimumAbsDifferenceBST(n00) == 5)

	n00.Left = &TreeNode{Val: 3}
	fmt.Println(minimumAbsDifferenceBST(n00) == 2)

	n00 = &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}
	fmt.Println(minimumAbsDifferenceBST(n00) == 2)

	n00 = &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 6}}
	fmt.Println(minimumAbsDifferenceBST(n00) == 1)

	n00 = &TreeNode{Val: 20,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 22}, Right: &TreeNode{Val: 99}}}
	fmt.Println(minimumAbsDifferenceBST(n00) == 2)
}

func minimumAbsDifferenceBST(root *TreeNode) int {
	return getMinimumAbsDifferenceBST(root, nil, nil)
}

func getMinimumAbsDifferenceBST(root *TreeNode, l, r *int) int {
	a := abs(root.Val, l, r)

	if root.Left != nil {
		r := getMinimumAbsDifferenceBST(root.Left, l, &root.Val)
		if r < a {
			a = r
		}
	}
	if root.Right != nil {
		r := getMinimumAbsDifferenceBST(root.Right, &root.Val, r)
		if r < a {
			a = r
		}
	}
	return a
}

func abs(a int, b, c *int) int {
	r := math.MaxInt
	if b != nil {
		r = a - *b
	}
	if c != nil && r > *c-a {
		r = *c - a
	}
	return r
}

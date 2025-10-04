package main

import (
	"fmt"
	"math"
)

func testMaximumDepthOfBinaryTree() {
	fmt.Println(MaximumDepthOfBinaryTree(nil, 0) == 0)

	n00 := &TreeNode{}
	fmt.Println(MaximumDepthOfBinaryTree(n00, 0) == 1)

	n10 := &TreeNode{}
	n00.Left = n10
	fmt.Println(MaximumDepthOfBinaryTree(n00, 0) == 2)

	n20 := &TreeNode{}
	n10.Left = n20
	fmt.Println(MaximumDepthOfBinaryTree(n00, 0) == 3)

	n11 := &TreeNode{}
	n00.Right = n11
	fmt.Println(MaximumDepthOfBinaryTree(n00, 0) == 3)
}

func MaximumDepthOfBinaryTree(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}
	dep += 1
	ans := dep
	if root.Left != nil {
		ans = int(math.Max(float64(MaximumDepthOfBinaryTree(root.Left, dep)), float64(ans)))
	}
	if root.Right != nil {
		ans = int(math.Max(float64(MaximumDepthOfBinaryTree(root.Right, dep)), float64(ans)))
	}
	return ans
}

func testMaximumDepthOfBinaryTree2() {
	s := 0
	MaximumDepthOfBinaryTree2(nil, 0, &s)
	fmt.Println(s == 0)

	n00 := &TreeNode{}
	s = 0
	MaximumDepthOfBinaryTree2(n00, 0, &s)
	fmt.Println(s == 1)

	n10 := &TreeNode{}
	n00.Left = n10
	s = 0
	MaximumDepthOfBinaryTree2(n00, 0, &s)
	fmt.Println(s == 2)

	n20 := &TreeNode{}
	n10.Left = n20
	s = 0
	MaximumDepthOfBinaryTree2(n00, 0, &s)
	fmt.Println(s == 3)

	n11 := &TreeNode{}
	n00.Right = n11
	s = 0
	MaximumDepthOfBinaryTree2(n00, 0, &s)
	fmt.Println(s == 3)
}

func MaximumDepthOfBinaryTree2(root *TreeNode, dep int, ans *int) {
	if root == nil {
		return
	}
	dep += 1
	if dep > *ans {
		*ans = dep
	}
	if root.Left != nil {
		MaximumDepthOfBinaryTree2(root.Left, dep, ans)
	}
	if root.Right != nil {
		MaximumDepthOfBinaryTree2(root.Right, dep, ans)
	}
}

func testMaximumDepthOfBinaryTree3() {
	fmt.Println(MaximumDepthOfBinaryTree3(nil) == 0)

	n00 := &TreeNode{}
	fmt.Println(MaximumDepthOfBinaryTree3(n00) == 1)

	n10 := &TreeNode{}
	n00.Left = n10
	fmt.Println(MaximumDepthOfBinaryTree3(n00) == 2)

	n20 := &TreeNode{}
	n10.Left = n20
	fmt.Println(MaximumDepthOfBinaryTree3(n00) == 3)

	n11 := &TreeNode{}
	n00.Right = n11
	fmt.Println(MaximumDepthOfBinaryTree3(n00) == 3)
}

func MaximumDepthOfBinaryTree3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left_ans, right_ans := 0, 0
	if root.Left != nil {
		left_ans = MaximumDepthOfBinaryTree3(root.Left)
	}
	if root.Right != nil {
		right_ans = MaximumDepthOfBinaryTree3(root.Right)
	}
	ans := 1
	if left_ans > right_ans {
		ans += left_ans
	} else {
		ans += right_ans
	}
	return ans
}

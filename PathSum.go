package main

import "fmt"

func TestPathSum() {
	fmt.Println(pathSum(nil, 0) == false)

	n00 := &TreeNode{Val: 0}
	fmt.Println(pathSum(n00, 0) == true)

	n10 := &TreeNode{Val: 1}
	n00.Left = n10
	fmt.Println(pathSum(n00, 2) == false)
	fmt.Println(pathSum(n00, 1) == true)

	n11 := &TreeNode{Val: 2}
	n00.Right = n11
	fmt.Println(pathSum(n00, 2) == true)
	fmt.Println(pathSum(n00, 3) == false)

	n20 := &TreeNode{Val: 2}
	n10.Left = n20
	fmt.Println(pathSum(n00, 3) == true)
}

func pathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if (sum - root.Val) == 0 {
			return true
		}
		return false
	}

	if root.Left != nil {
		if pathSum(root.Left, sum-root.Val) {
			return true
		}
	}

	if root.Right != nil {
		if pathSum(root.Right, sum-root.Val) {
			return true
		}
	}
	return false
}

func TestPathSum2() {
	fmt.Println(pathSum2(nil, 0) == false)

	n00 := &TreeNode{Val: 0}
	fmt.Println(pathSum2(n00, 0) == true)

	n10 := &TreeNode{Val: 1}
	n00.Left = n10
	fmt.Println(pathSum2(n00, 2) == false)
	fmt.Println(pathSum2(n00, 1) == true)

	n11 := &TreeNode{Val: 2}
	n00.Right = n11
	fmt.Println(pathSum2(n00, 2) == true)
	fmt.Println(pathSum2(n00, 3) == false)

	n20 := &TreeNode{Val: 2}
	n10.Left = n20
	fmt.Println(pathSum2(n00, 3) == true)

	n00.Val = 5
	n10.Val = 4
	n11.Val = 8
	n20.Val = 11
	n22 := &TreeNode{Val: 13}
	n11.Left = n22
	n23 := &TreeNode{Val: 4}
	n11.Right = n23
	n30 := &TreeNode{Val: 7}
	n20.Left = n30
	n31 := &TreeNode{Val: 2}
	n20.Right = n31
	n37 := &TreeNode{Val: 1}
	n23.Right = n37
	fmt.Println(pathSum2(n00, 22) == true)
}

func pathSum2(root *TreeNode, sum int) bool {
	ans := pathSum2It(root)
	fmt.Println(ans)
	for _, v := range ans {
		if v == sum {
			return true
		}
	}
	return false
}

func pathSum2It(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	ans := []int{}
	if root.Left != nil {
		ans = append(ans, pathSum2It(root.Left)...)
	}

	if root.Right != nil {
		ans = append(ans, pathSum2It(root.Right)...)
	}
	for i, _ := range ans {
		ans[i] += root.Val
	}
	return ans
}

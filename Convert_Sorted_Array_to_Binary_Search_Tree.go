package main

import "fmt"

func testConvertSortedArray2BinarySearchTree() {
	n := ConvertSortedArray2BinarySearchTree([]int{})
	fmt.Println("0", n == nil)

	n = ConvertSortedArray2BinarySearchTree([]int{-10, -3, 0, 5, 9})
	fmt.Println("1", n.Val == 0)
	fmt.Println("1", n.Left.Val == -3)
	fmt.Println("1", n.Left.Left.Val == -10)
	fmt.Println("1", n.Right.Val == 9)
	fmt.Println("1", n.Right.Left.Val == 5)

	n = ConvertSortedArray2BinarySearchTree([]int{1, 3})
	fmt.Println("2", n.Val == 3)
	fmt.Println("2", n.Left.Val == 1)

	n = ConvertSortedArray2BinarySearchTree([]int{1})
	fmt.Println("3", n.Val == 1)
}

func ConvertSortedArray2BinarySearchTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	rc := &TreeNode{Val: nums[mid]}
	if mid > 0 {
		rc.Left = ConvertSortedArray2BinarySearchTree(nums[:mid])
	}
	if mid < len(nums)-1 {
		rc.Right = ConvertSortedArray2BinarySearchTree(nums[mid+1:])
	}
	return rc
}

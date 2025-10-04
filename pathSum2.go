package main

import (
	"container/list"
	"fmt"
	"slices"
)

func TestPathSumII() {
	fmt.Println(len(pathSumII(nil, 0)) == 0)

	n00 := &TreeNode{Val: 0}
	r := pathSumII(n00, 0)
	fmt.Println(len(r) == 1)
	fmt.Println(slices.Equal(r[0], []int{0}))

	n10 := &TreeNode{Val: 1}
	n00.Left = n10
	r = pathSumII(n00, 0)
	fmt.Println(len(r) == 0)
	r = pathSumII(n00, 2)
	fmt.Println(len(r) == 0)
	r = pathSumII(n00, 1)
	fmt.Println(len(r) == 1)
	fmt.Println(slices.Equal(r[0], []int{0, 1}))

	n11 := &TreeNode{Val: 2}
	n00.Right = n11
	r = pathSumII(n00, 2)
	fmt.Println(len(r) == 1)
	fmt.Println(slices.Equal(r[0], []int{0, 2}))

	n20 := &TreeNode{Val: 1}
	n10.Left = n20
	r = pathSumII(n00, 2)
	fmt.Println(len(r) == 2)
	fmt.Println(slices.Equal(r[0], []int{0, 1, 1}))
	fmt.Println(slices.Equal(r[1], []int{0, 2}))
}

var cache = list.List{}
var rc = [][]int{}

func pathSumII(root *TreeNode, sum int) [][]int {
	rc = make([][]int, 0)
	getSumPath(root, sum)
	return rc
}

// top-down, preorder
func getSumPath(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	//root
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			cache.PushBack(root.Val)
			rc = append(rc, list2Array(cache))
			cache.Remove(cache.Back())
			return
		}
	}

	//left
	if root.Left != nil {
		cache.PushBack(root.Val)
		getSumPath(root.Left, sum-root.Val)
		cache.Remove(cache.Back())
	}
	//right
	if root.Right != nil {
		cache.PushBack(root.Val)
		getSumPath(root.Right, sum-root.Val)
		cache.Remove(cache.Back())
	}
}

func list2Array(cache list.List) []int {
	r := make([]int, 0)
	for i := cache.Front(); i != nil; i = i.Next() {
		r = append(r, i.Value.(int))
	}

	return r
}

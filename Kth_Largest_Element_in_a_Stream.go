package main

import (
	"container/list"
	"fmt"
)

func testKthLargest() {
	k := ConstructorKthLargest(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3) == 4)
	fmt.Println(k.Add(5) == 5)
	fmt.Println(k.Add(10) == 5)
	fmt.Println(k.Add(9) == 8)
	fmt.Println(k.Add(4) == 8)

	k = ConstructorKthLargest(4, []int{7, 7, 7, 8, 3})
	fmt.Println(k.Add(2) == 7)
	fmt.Println(k.Add(10) == 7)
	fmt.Println(k.Add(9) == 7)
	fmt.Println(k.Add(9) == 8)
}

type KthLargest struct {
	root  *TreeNode
	count int
	cap   int
}

func ConstructorKthLargest(k int, nums []int) KthLargest {
	if k <= 0 {
		panic("invalid capability")
	}
	kth := KthLargest{cap: k}
	for i := 0; i < len(nums); i++ {
		kth.Add(nums[i])
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	s := getTheSmallest(this.root)
	if this.count < this.cap || *s < val {
		this.root = insertBSTWithSameValue(this.root, val)
		this.count++
	}

	if this.count > this.cap {
		this.root = deleteTheSmallest(this.root)
		this.count--
	}

	return *getTheSmallest(this.root)
}

func getMaxBSTDiff(root *TreeNode) *int {
	s := getTheSmallest(root)
	m := getTheMaximum(root)
	if s != nil && m != nil {
		t := *m - *s
		return &t
	}
	return nil
}

func getSmallBSTDiff(root *TreeNode) *int {
	if root == nil {
		return nil
	}
	cache := list.New()
	n := root
	for n != nil {
		cache.PushBack(n)
		n = n.Left
	}

	var lastValue, minimum *int
	for cache.Back() != nil {
		cur := cache.Back().Value.(*TreeNode)
		cache.Remove(cache.Back())
		// process cur
		if lastValue != nil {
			t := cur.Val - *lastValue
			if minimum == nil || t < *minimum {
				minimum = &t
			}
		}
		lastValue = &cur.Val

		if cur.Right != nil {
			n = cur.Right
			for n != nil {
				cache.PushBack(n)
				n = n.Left
			}
		}
	}
	return minimum
}

func getTheSmallest(root *TreeNode) *int {
	if root == nil {
		return nil
	}
	n := root
	for n.Left != nil {
		n = n.Left
	}
	return &n.Val
}

func getTheMaximum(root *TreeNode) *int {
	if root == nil {
		return nil
	}
	n := root
	for n.Right != nil {
		n = n.Right
	}
	return &n.Val
}

func deleteTheSmallest(root *TreeNode) *TreeNode {
	if root == nil {
		panic("delete nil element")
	}
	var pn *TreeNode
	n := root
	for n.Left != nil {
		pn = n
		n = n.Left
	}
	if n == root {
		return root.Right
	}
	pn.Left = n.Right
	return root
}

func insertBSTWithSameValue(root *TreeNode, val int) *TreeNode {
	n := searchBSTNodeWithSameValue(root, val)
	if n == nil {
		return &TreeNode{Val: val}
	}
	if n.Val == val {
		node := &TreeNode{Val: val, Left: n.Left}
		n.Left = node
	}
	if n.Val < val {
		n.Right = &TreeNode{Val: val}
	}
	if n.Val > val {
		n.Left = &TreeNode{Val: val}
	}
	return root
}

func searchBSTNodeWithSameValue(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		if root.Right != nil {
			return searchBSTNodeWithSameValue(root.Right, val)
		}
		return root
	}
	if root.Val > val {
		if root.Left != nil {
			return searchBSTNodeWithSameValue(root.Left, val)
		}
		return root
	}
	return nil
}

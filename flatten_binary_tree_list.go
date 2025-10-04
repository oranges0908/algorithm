package main

import "C"
import (
	"container/list"
	"fmt"
)

func TestFlattenBinaryTree2List() {

	fmt.Println(CompareTree(FlattenBinaryTree2List(nil), nil))
	n00 := &TreeNode{Val: 0}
	t00 := &TreeNode{Val: 0}
	fmt.Println(CompareTree(FlattenBinaryTree2List(n00), t00))
	n11 := &TreeNode{Val: 11}
	n00.Right = n11
	t11 := &TreeNode{Val: 11}
	t00.Right = t11
	fmt.Println(CompareTree(FlattenBinaryTree2List(n00), t00))
	n22 := &TreeNode{Val: 22}
	n11.Left = n22
	t23 := &TreeNode{Val: 22}
	t11.Right = t23
	fmt.Println(CompareTree(FlattenBinaryTree2List(n00), t00))
	n23 := &TreeNode{Val: 23}
	n11.Right = n23
	t37 := &TreeNode{Val: 23}
	t23.Right = t37

	//fmt.Println(tree2stringPreorder(n00))
	//fmt.Println(tree2stringInorder(n00))
	//fmt.Println(tree2stringPostorder(n00))
	r := FlattenBinaryTree2List(n00)
	fmt.Println(t2s(r))
	fmt.Println(CompareTree(r, t00))
}

// var cache = list.List{}
func tree2stringInorder(root *TreeNode) []string {
	rc := make([]string, 0)
	cache = list.List{}
	curr := root
	for {
		if curr != nil {
			cache.PushBack(curr)
			curr = curr.Left
			continue
		}

		b := cache.Back()
		cache.Remove(b)
		curr = b.Value.(*TreeNode)
		rc = append(rc, fmt.Sprint(curr.Val))

		curr = curr.Right

		if curr == nil && cache.Len() == 0 {
			break
		}
	}

	return rc
}

func tree2stringPreorder(root *TreeNode) []string {
	rc := make([]string, 0)
	cache = list.List{}
	cache.PushBack(root)
	for {
		b := cache.Back()
		cache.Remove(b)
		curr := b.Value.(*TreeNode)
		rc = append(rc, fmt.Sprint(curr.Val))

		if curr.Right != nil {
			cache.PushBack(curr.Right)
		}

		if curr.Left != nil {
			cache.PushBack(curr.Left)
		}

		if cache.Len() == 0 {
			break
		}
	}

	return rc
}

func tree2stringPostorder(root *TreeNode) []string {
	rc := make([]string, 0)
	cache = list.List{}
	curr := root
	lastV := &TreeNode{}
	for {
		if curr != nil {
			cache.PushBack(curr)
			curr = curr.Left
			continue
		}
		b := cache.Back()
		curr = b.Value.(*TreeNode)
		if curr.Right != nil && curr.Right != lastV {
			curr = curr.Right
			continue
		}
		lastV = curr
		//fmt.Println(treeList2Array(cache))
		rc = append(rc, fmt.Sprint(curr.Val))
		cache.Remove(b)
		curr = nil
		if cache.Len() == 0 {
			break
		}
	}

	return rc
}

func treeList2Array(cache list.List) []int {
	r := make([]int, 0)
	for i := cache.Front(); i != nil; i = i.Next() {
		r = append(r, i.Value.(*TreeNode).Val)
	}

	return r
}

func FlattenBinaryTree2List(root *TreeNode) *TreeNode {
	Head := &TreeNode{}
	ProcessFlattenBinaryTree2List(root, Head)
	return Head.Right
}

func ProcessFlattenBinaryTree2List(root *TreeNode, list *TreeNode) {
	// pre-order
	// val
	if root == nil {
		return
	}
	list.Right = &TreeNode{Val: root.Val}
	// left
	if root.Left != nil {
		ProcessFlattenBinaryTree2List(root.Left, list.Right)
		list = list.Right
	}

	// right
	if root.Right != nil {
		ProcessFlattenBinaryTree2List(root.Right, list.Right)
		list = list.Right
	}
}

func t2s(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	curr := root
	rc := make([]string, 0)
	for {
		if curr == nil {
			break
		}
		rc = append(rc, fmt.Sprint(curr.Val))
		if curr.Right != nil {
			rc = append(rc, "null")
			curr = curr.Right
		} else {
			curr = nil
		}
	}
	return rc
}

package main

import (
	"container/list"
	"fmt"
)

func testMinDepth() {
	fmt.Println(minDepth(nil) == 0)
	tree1 := &TreeNode{Val: 3}
	fmt.Println(minDepth(tree1) == 1)
	tree2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}}
	fmt.Println(minDepth(tree2) == 2)
	tree3 := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}
	fmt.Println(minDepth(tree3) == 2)
	tree4 := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(minDepth(tree4) == 2)
	tree5 := &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}
	fmt.Println(minDepth(tree5) == 5)
	tree6 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	fmt.Println(minDepth(tree6) == 2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	toProcessList1 := list.New()
	toProcessList2 := list.New()

	toProcessList := toProcessList1
	cacheList := toProcessList2
	toProcessList.PushBack(root)
	deep := 0
loop:
	for {
		deep++
		for toProcessList.Len() > 0 {
			tp := toProcessList.Front()
			node := tp.Value.(*TreeNode)
			if node.Left == nil && node.Right == nil {
				break loop
			}
			if node.Left != nil {
				cacheList.PushBack(node.Left)
			}
			if node.Right != nil {
				cacheList.PushBack(node.Right)
			}
			toProcessList.Remove(tp)
		}
		if deep%2 == 0 {
			toProcessList = toProcessList1
			cacheList = toProcessList2
		} else {
			toProcessList = toProcessList2
			cacheList = toProcessList1
		}
	}
	return deep
}

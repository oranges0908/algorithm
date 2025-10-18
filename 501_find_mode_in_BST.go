package main

import (
	"container/list"
	"fmt"
	"math"
	"reflect"
)

func TestFindModeInBST() {

	n00 := &TreeNode{Val: 5}
	fmt.Println(reflect.DeepEqual(findModeInBST(n00), []int{5}))

	n00.Left = &TreeNode{Val: 1}
	fmt.Println(reflect.DeepEqual(findModeInBST(n00), []int{1, 5}))

	n00.Right = &TreeNode{Val: 10}
	fmt.Println(reflect.DeepEqual(findModeInBST(n00), []int{1, 5, 10}))

	n00.Left = &TreeNode{Val: 5}
	fmt.Println(reflect.DeepEqual(findModeInBST(n00), []int{5}))

	n00 = &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 7},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 20}}}
	fmt.Println(reflect.DeepEqual(findModeInBST(n00), []int{7, 15}))

	n00 = &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 7},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 15}}}
	fmt.Println(reflect.DeepEqual(findModeInBST(n00), []int{15}))
}

func findModeInBST(root *TreeNode) []int {
	rc := make([]int, 0)
	lastValue := math.MaxInt
	maxCount := 0
	count := 0
	//preorder iteration
	l := list.New()
	cur := root
	for {
		if cur != nil {
			l.PushBack(cur)
			cur = cur.Left
			continue
		}

		if l.Len() == 0 {
			break
		}

		cur = l.Back().Value.(*TreeNode)
		l.Remove(l.Back())

		if cur.Val == lastValue {
			count++
		} else {
			count = 1
		}
		if count == maxCount {
			rc = append(rc, cur.Val)
		}
		if count > maxCount {
			rc = []int{cur.Val}
			maxCount = count
		}
		lastValue = cur.Val

		cur = cur.Right
	}
	return rc
}

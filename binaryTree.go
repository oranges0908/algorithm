package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type BTNode struct {
	Val   int
	Left  *BTNode
	Right *BTNode
}

func testBTPreorder() {
	testBT0()
	testBT1()
	testBT2()
	testBT3()
}

func testBT0() {
	//	fmt.Println(reflect.DeepEqual(bt_preorder_recursion(nil), []int{}))
	//	fmt.Println(reflect.DeepEqual(bt_preorder_iteration(nil), []int{}))
	//fmt.Println(reflect.DeepEqual(bt_inorder_recursion(nil), []int{}))
	//fmt.Println(reflect.DeepEqual(bt_inorder_iteration(nil), []int{}))
	//fmt.Println(reflect.DeepEqual(bt_postorder_recursion(nil), []int{}))
	//fmt.Println(reflect.DeepEqual(bt_postorder_iteration(nil), []int{}))
	fmt.Println(reflect.DeepEqual(bt_levelorder_recursion(nil), []int{}))
	fmt.Println(reflect.DeepEqual(bt_levelorder_iteration(nil), []int{}))
}

func testBT1() {
	n0 := &BTNode{1, nil, nil}
	n10 := &BTNode{2, nil, nil}
	n20 := &BTNode{3, nil, nil}

	n0.Right = n10
	n10.Left = n20
	//rc := []int{1, 2, 3}
	//rc := []int{1, 3, 2}
	//rc := []int{3, 2, 1}
	rc := []int{1, 2, 3}

	//rcc := bt_preorder_recursion(n0)
	//rcc := bt_inorder_recursion(n0)
	//rcc := bt_postorder_recursion(n0)
	rcc := bt_levelorder_recursion(n0)
	fmt.Println(rcc)
	fmt.Println(reflect.DeepEqual(rcc, rc))
	//rci := bt_preorder_iteration(n0)
	//rci := bt_inorder_iteration(n0)
	//rci := bt_postorder_iteration(n0)
	rci := bt_levelorder_iteration(n0)
	fmt.Println(rci)
	fmt.Println(reflect.DeepEqual(rci, rc))
}

func testBT2() {
	n0 := &BTNode{1, nil, nil}
	n10 := &BTNode{2, nil, nil}
	n11 := &BTNode{3, nil, nil}
	n20 := &BTNode{4, nil, nil}
	n21 := &BTNode{5, nil, nil}
	n22 := &BTNode{8, nil, nil}
	n30 := &BTNode{6, nil, nil}
	n31 := &BTNode{7, nil, nil}
	n32 := &BTNode{9, nil, nil}

	n0.Left = n10
	n0.Right = n11
	n10.Left = n20
	n10.Right = n21
	n11.Right = n22
	n21.Left = n30
	n21.Right = n31
	n22.Left = n32

	//rc := []int{1, 2, 4, 5, 6, 7, 3, 8, 9}
	//rc := []int{4, 2, 6, 5, 7, 1, 3, 9, 8}
	//rc := []int{4, 6, 7, 5, 2, 9, 8, 3, 1}
	rc := []int{1, 2, 3, 4, 5, 8, 6, 7, 9}

	//rcc := bt_preorder_recursion(n0)
	//rcc := bt_inorder_recursion(n0)
	//rcc := bt_postorder_recursion(n0)
	rcc := bt_levelorder_recursion(n0)
	fmt.Println(rcc)
	fmt.Println(reflect.DeepEqual(rcc, rc))
	//rci := bt_preorder_iteration(n0)
	//rci := bt_inorder_iteration(n0)
	//rci := bt_postorder_iteration(n0)
	rci := bt_levelorder_iteration(n0)
	fmt.Println(rci)
	fmt.Println(reflect.DeepEqual(rci, rc))
}

func testBT3() {
	n0 := &BTNode{1, nil, nil}
	rc := []int{1}

	//rcc := bt_preorder_recursion(n0)
	//rcc := bt_inorder_recursion(n0)
	//rcc := bt_postorder_recursion(n0)
	rcc := bt_levelorder_recursion(n0)
	fmt.Println(rcc)
	fmt.Println(reflect.DeepEqual(rcc, rc))
	//rci := bt_preorder_iteration(n0)
	//rci := bt_inorder_iteration(n0)
	//rci := bt_postorder_iteration(n0)
	rci := bt_levelorder_iteration(n0)
	fmt.Println(rci)
	fmt.Println(reflect.DeepEqual(rci, rc))
}

func bt_preorder_recursion(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{node.Val}
	if node.Left != nil {
		rc = append(rc, bt_preorder_recursion(node.Left)...)
	}
	if node.Right != nil {
		rc = append(rc, bt_preorder_recursion(node.Right)...)
	}
	return rc
}

func bt_preorder_iteration(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	c := list.List{}
	c.PushBack(node)
	for d := c.Back(); d != nil; d = c.Back() {
		n := d.Value.(*BTNode)
		// 前序遍历
		rc = append(rc, n.Val)
		if n.Right != nil {
			c.PushBack(n.Right)
		}
		if n.Left != nil {
			c.PushBack(n.Left)
		}
		// 弹出
		c.Remove(d)
	}

	return rc
}

func bt_inorder_recursion(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	if node.Left != nil {
		rc = append(rc, bt_inorder_recursion(node.Left)...)
	}
	rc = append(rc, node.Val)
	if node.Right != nil {
		rc = append(rc, bt_inorder_recursion(node.Right)...)
	}
	return rc
}

func bt_inorder_iteration(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	c := list.List{}
	curr := node
	for {
		// 左侧走到底
		if curr != nil {
			c.PushBack(curr)
			curr = curr.Left
			continue
		}

		curr = c.Back().Value.(*BTNode)
		c.Remove(c.Back()) // 弹出

		rc = append(rc, curr.Val)

		curr = curr.Right

		if c.Len() == 0 && curr == nil {
			break
		}
	}

	return rc
}

func bt_postorder_recursion(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	if node.Left != nil {
		rc = append(rc, bt_postorder_recursion(node.Left)...)
	}
	if node.Right != nil {
		rc = append(rc, bt_postorder_recursion(node.Right)...)
	}
	rc = append(rc, node.Val)
	return rc
}

func bt_postorder_iteration(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	c := list.List{}
	curr := node
	var lastVisit *BTNode
	for {
		if curr != nil {
			c.PushBack(curr)
			curr = curr.Left
			continue
		}

		curr = c.Back().Value.(*BTNode)
		if curr.Right != nil && curr.Right != lastVisit {
			curr = curr.Right
			continue
		}

		rc = append(rc, curr.Val)
		lastVisit = curr
		c.Remove(c.Back())
		curr = nil

		if c.Len() == 0 {
			break
		}
	}

	return rc
}

func bt_levelorder_recursion(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	return bt_levelorder_recursion_p([]*BTNode{node})
}

func bt_levelorder_recursion_p(nodes []*BTNode) []int {
	if len(nodes) == 0 {
		return []int{}
	}
	rc := []int{}
	nextLevel := []*BTNode{}
	for _, node := range nodes {
		rc = append(rc, node.Val)
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}

	return append(rc, bt_levelorder_recursion_p(nextLevel)...)
}

func bt_levelorder_iteration(node *BTNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	c := list.List{}
	c.PushBack(node)

	for {
		d := c.Front()
		if d == nil {
			break
		}
		n := d.Value.(*BTNode)
		rc = append(rc, n.Val)
		c.Remove(d)
		if n.Left != nil {
			c.PushBack(n.Left)
		}
		if n.Right != nil {
			c.PushBack(n.Right)
		}
	}

	return rc
}

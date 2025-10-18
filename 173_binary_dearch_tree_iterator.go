package main

import (
	"container/list"
	"fmt"
)

func testBSTIterator() {
	b0 := ConstructorBSTIterator(nil)
	fmt.Println(b0.HasNext() == false)

	n00 := &TreeNode{Val: 0}
	b1 := ConstructorBSTIterator(n00)
	fmt.Println(b1.HasNext() == true)
	fmt.Println(b1.Next() == 0)
	fmt.Println(b1.HasNext() == false)

	n11 := &TreeNode{Val: 10}
	n00.Right = n11
	b2 := ConstructorBSTIterator(n00)
	fmt.Println(b2.HasNext() == true)
	fmt.Println(b2.Next() == 0)
	fmt.Println(b2.HasNext() == true)
	fmt.Println(b2.Next() == 10)
	fmt.Println(b2.HasNext() == false)

	n00 = &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}
	b3 := ConstructorBSTIterator(n00)
	fmt.Println(b3.Next() == 3)
	//fmt.Println(b3.Next() == 7)
	fmt.Println(b3.Next())
	fmt.Println(b3.HasNext() == true)
	fmt.Println(b3.Next() == 9)
	fmt.Println(b3.HasNext() == true)
	fmt.Println(b3.Next() == 15)
	fmt.Println(b3.HasNext() == true)
	fmt.Println(b3.Next() == 20)
	fmt.Println(b3.HasNext() == false)
}

type BSTIterator struct {
	cache *list.List
}

func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	l := list.List{}
	n := root
	// 左侧全部入栈
	for n != nil {
		l.PushBack(n)
		n = n.Left
	}
	return BSTIterator{cache: &l}
}

func (this *BSTIterator) Next() int {
	// 中序遍历
	h := this.cache.Back().Value.(*TreeNode)
	this.cache.Remove(this.cache.Back())
	// 右侧待处理
	if h.Right != nil {
		n := h.Right
		// 左侧全部入栈
		for n != nil {
			this.cache.PushBack(n)
			n = n.Left
		}
	}
	return h.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.cache.Len() > 0
}

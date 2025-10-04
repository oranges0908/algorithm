package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
	next  *Node
}

func TestSetNextRightPointer() {
	n00 := &Node{val: 0}
	n10 := &Node{val: 10}
	n00.left = n10
	n11 := &Node{val: 11}
	n00.right = n11
	n20 := &Node{val: 20}
	n10.left = n20
	n21 := &Node{val: 21}
	n10.right = n21
	n22 := &Node{val: 22}
	n11.left = n22
	n23 := &Node{val: 23}
	n11.right = n23

	SetNextRightPointer(n00)
	fmt.Println(n00.next == nil)
	fmt.Println(n10.next == n11)
	fmt.Println(n11.next == nil)
	fmt.Println(n20.next == n21)
	fmt.Println(n21.next == n22)
	fmt.Println(n22.next == n23)
	fmt.Println(n23.next == nil)
}

func SetNextRightPointer(n *Node) {
	cache := list.List{}
	ncache := list.List{}
	cache.PushBack(n)
	var pre *Node
	for {
		for e := cache.Front(); e != nil; e = e.Next() {
			node := e.Value.(*Node)
			if pre != nil {
				pre.next = node
			}
			pre = node
			if node.left != nil {
				ncache.PushBack(node.left)
			}
			if node.right != nil {
				ncache.PushBack(node.right)
			}
		}
		cache = ncache
		ncache = list.List{}
		pre = nil
		if cache.Len() == 0 {
			break
		}
	}
}

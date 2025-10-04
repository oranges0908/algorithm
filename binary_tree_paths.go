package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func TestAllPath() {
	fmt.Println(len(allPath(nil)) == 0)

	n00 := &TreeNode{Val: 0}
	r := allPath(n00)
	fmt.Println(len(r) == 1)
	fmt.Println(r[0] == "0")

	n10 := &TreeNode{Val: 1}
	n00.Left = n10
	r = allPath(n00)
	fmt.Println(len(r) == 1)
	fmt.Println(r[0] == "0->1")

	n11 := &TreeNode{Val: 2}
	n00.Right = n11
	r = allPath(n00)
	fmt.Println(len(r) == 2)
	fmt.Println(r[0] == "0->1")
	fmt.Println(r[1] == "0->2")

	n20 := &TreeNode{Val: 1}
	n10.Left = n20
	r = allPath(n00)
	fmt.Println(len(r) == 2)
	fmt.Println(r[0] == "0->1->1")
	fmt.Println(r[1] == "0->2")
}

var rc_01 = []string{}

//var cache = list.List{}

func allPath(root *TreeNode) []string {
	rc_01 = make([]string, 0)
	cache = list.List{}
	getAllPath(root)
	return rc_01
}

// top-down, preorder
func getAllPath(root *TreeNode) {
	if root == nil {
		return
	}
	//root
	if root.Left == nil && root.Right == nil {
		cache.PushBack(root.Val)
		rc_01 = append(rc_01, list2string(cache))
		cache.Remove(cache.Back())
		return
	}

	//left
	if root.Left != nil {
		cache.PushBack(root.Val)
		getAllPath(root.Left)
		cache.Remove(cache.Back())
	}
	//right
	if root.Right != nil {
		cache.PushBack(root.Val)
		getAllPath(root.Right)
		cache.Remove(cache.Back())
	}
}

func list2string(cache list.List) string {
	r := ""
	for i := cache.Front(); i != nil; i = i.Next() {
		r += strconv.Itoa(i.Value.(int))
		r += "->"
	}
	if len(r) > 2 {
		r = r[0 : len(r)-2]
	}
	return r
}

package main

import (
	"fmt"
	"math"
)

func TestConstructBSTFromPreorder() {
	n := constructBSTFromPreorder([]int{})
	fmt.Println("0", n == nil)

	n = constructBSTFromPreorder([]int{1})
	fmt.Println("3", n.Val == 1)

	n = constructBSTFromPreorder([]int{1, 3})
	fmt.Println("2", n.Val == 1)
	fmt.Println("2", n.Right.Val == 3)

	n = constructBSTFromPreorder([]int{8, 5, 1, 7, 10, 12})
	fmt.Println("1", n.Val == 8)
	fmt.Println("1", n.Left.Val == 5)
	fmt.Println("1", n.Left.Left.Val == 1)
	fmt.Println("1", n.Left.Right.Val == 7)
	fmt.Println("1", n.Right.Val == 10)
	fmt.Println("1", n.Right.Right.Val == 12)
}

var cur int

func constructBSTFromPreorder(preorder []int) *TreeNode {
	cur = 0
	return preorder2BST(preorder, math.MinInt, math.MaxInt)
}

func preorder2BST(preorder []int, lower, upper int) *TreeNode {
	if len(preorder) == cur {
		return nil
	}

	if preorder[cur] < lower || preorder[cur] > upper {
		return nil
	}
	v := preorder[cur]
	cur++
	r := &TreeNode{Val: v}
	r.Left = preorder2BST(preorder, lower, v)
	r.Right = preorder2BST(preorder, v, upper)
	return r
}

func constructBSTFromPreorder1(preorder []int) *TreeNode {
	ng := make([]int, len(preorder))

	maxStack := make([]int, 0)
	for i := len(preorder) - 1; i >= 0; i-- {
		for len(maxStack) > 0 && preorder[i] > preorder[maxStack[len(maxStack)-1]] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		if len(maxStack) > 0 {
			ng[i] = maxStack[len(maxStack)-1]
		} else {
			ng[i] = -1
		}
		maxStack = append(maxStack, i)
	}
	return preorder2BST1(preorder, ng, 0)
}

func preorder2BST1(preorder, nextGreater []int, base int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//前序遍历
	//root
	r := &TreeNode{Val: preorder[0]}

	//left
	lmt := nextGreater[0] - base
	if lmt < 0 {
		lmt = len(nextGreater)
	}
	if len(preorder) > 1 {
		r.Left = preorder2BST1(preorder[1:lmt], nextGreater[1:lmt], base+1)
	}

	//right
	if lmt < len(nextGreater) {
		r.Right = preorder2BST1(preorder[lmt:], nextGreater[lmt:], base+lmt)
	}
	return r
}

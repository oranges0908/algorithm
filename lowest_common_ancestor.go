package main

import "fmt"

func TestLowestCommonAncestor() {
	n00 := &TreeNode{Val: 0}
	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	n11 := &TreeNode{Val: 11}
	n00.Right = n11
	n20 := &TreeNode{Val: 20}
	n10.Left = n20
	n21 := &TreeNode{Val: 21}
	n10.Right = n21
	n22 := &TreeNode{Val: 22}
	n11.Left = n22
	n23 := &TreeNode{Val: 23}
	n11.Right = n23

	//fmt.Println(LowestCommonAncestor(n00, &TreeNode{Val: 21}, &TreeNode{Val: 22}).Val == 0)
	//fmt.Println(LowestCommonAncestor(nil, &TreeNode{Val: 21}, &TreeNode{Val: 22}) == nil)
	//fmt.Println(LowestCommonAncestor(n00, nil, &TreeNode{Val: 22}) == nil)

	fmt.Println(lowestCommonAncestor(n00, &TreeNode{Val: 20}, &TreeNode{Val: 21}).Val == 10)
	fmt.Println(lowestCommonAncestor(n00, &TreeNode{Val: 21}, &TreeNode{Val: 22}).Val == 0)
	fmt.Println(lowestCommonAncestor(nil, &TreeNode{Val: 21}, &TreeNode{Val: 22}) == nil)
	n00 = &TreeNode{Val: 6,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}
	fmt.Println(lowestCommonAncestor(n00, &TreeNode{Val: 3}, &TreeNode{Val: 5}).Val == 4)
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil {
		return nil
	}
	pp := findNodePath(root, p)
	qp := findNodePath(root, q)
	var preVal *TreeNode
	for {
		if pp == nil || qp == nil {
			break
		}
		if pp.Val == qp.Val {
			preVal = pp
			pp = pp.Left
			qp = qp.Left
		} else {
			break
		}
	}
	if preVal != nil {
		return preVal.Right
	}
	return nil
}

func findNodePath(root, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	path := &TreeNode{Val: root.Val}
	if root.Val == p.Val {
		path.Right = root
		return path
	}
	if root.Left != nil {
		v := findNodePath(root.Left, p)
		if v != nil {
			path.Left = v
			path.Right = root
			return path
		}
	}
	if root.Right != nil {
		v := findNodePath(root.Right, p)
		if v != nil {
			path.Left = v
			path.Right = root
			return path
		}
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

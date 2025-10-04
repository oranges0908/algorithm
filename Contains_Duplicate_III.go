package main

import "fmt"

func testContainsNearbyAlmostDuplicate3() {
	fmt.Println(containsNearbyAlmostDuplicate3([]int{2, 2}, 2, 0) == true)
	fmt.Println(containsNearbyAlmostDuplicate3([]int{1, 2, 3, 1}, 3, 0) == true)
	fmt.Println(containsNearbyAlmostDuplicate3([]int{1, 5, 9, 1, 5, 9}, 2, 3) == false)
}

// 构建一个BST
// 不断的插入和移除
// 插入过程中，需要检查临近两个节点的Diff
func containsNearbyAlmostDuplicate3(nums []int, indexDiff int, valueDiff int) bool {
	root, _ := AddAndCheck(nil, nums[0], valueDiff)
	rc := false
	for i := 1; i < len(nums); i++ {
		if i > indexDiff {
			root = deleteBST(root, nums[i-indexDiff-1])
		}
		root, rc = AddAndCheck(root, nums[i], valueDiff)
		if rc {
			return true
		}
		if i%10 == 0 {
			root = balenceBST(root)
		}
	}
	return false
}

func findBSTAsideNode(root, leftP, rightP *TreeNode, val int) (*TreeNode, *TreeNode) {
	if root == nil {
		return leftP, rightP
	}
	if root.Val == val {
		return root.Left, root
	}
	if root.Val > val {
		return findBSTAsideNode(root.Left, leftP, root, val)
	}
	//root.Val < val
	return findBSTAsideNode(root.Right, root, rightP, val)
}

func AddAndCheck(root *TreeNode, val int, threshold int) (*TreeNode, bool) {
	lp, rp := findBSTAsideNode(root, nil, nil, val)
	if lp != nil && val-lp.Val <= threshold {
		return root, true
	}
	if rp != nil && rp.Val-val <= threshold {
		return root, true
	}
	if lp == nil && rp == nil {
		return &TreeNode{Val: val}, false
	}
	if lp == nil {
		rp.Left = &TreeNode{Val: val}
	}
	if rp == nil {
		lp.Right = &TreeNode{Val: val}
	}
	if lp != nil && rp != nil {
		t := &TreeNode{Val: val}
		if lp.Right != nil {
			t.Right = lp.Right
		}
		lp.Right = t
	}
	return root, false
}

func balenceBST(root *TreeNode) *TreeNode {
	ld := checkDepth(root.Left)
	rd := checkDepth(root.Right)
	m := 0
	if ld > rd {
		m = ld - rd
		for m > 1 {
			t := root.Left
			root.Left = t.Right
			t.Right = root
			root = t
			m--
		}
	} else {
		m = rd - ld
		for m > 1 {
			t := root.Right
			root.Right = t.Left
			t.Left = root
			root = t
			m--
		}
	}
	return root
}

func checkDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	t := 1
	if root.Left != nil {
		t = checkDepth(root.Left) + 1
	}
	if root.Right != nil {
		t1 := checkDepth(root.Right) + 1
		if t1 > t {
			t = t1
		}
	}
	return t
}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func TestCodec() {
	C := ConstructorST()
	fmt.Println(CompareTree(C.deserialize(C.serialize(nil)), nil))

	n00 := &TreeNode{Val: 0}
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n11 := &TreeNode{Val: 11}
	n00.Right = n11
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n23 := &TreeNode{Val: 23}
	n11.Right = n23
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n10 := &TreeNode{Val: 10}
	n00.Left = n10
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n20 := &TreeNode{Val: 20}
	n10.Left = n20
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))

}

func TestCodec1() {
	C := ConstructorST()
	n00 := &TreeNode{Val: 3}
	n10 := &TreeNode{Val: 2}
	n00.Left = n10
	n11 := &TreeNode{Val: 4}
	n00.Right = n11
	n20 := &TreeNode{Val: 3}
	n10.Left = n20
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
}

type Codec struct {
}

func ConstructorST() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	preorder := bt_tree_preorder_recursion(root)
	inorder := bt_tree_inorder_recursion(root)

	//buf := make([]byte, unsafe.Sizeof(preorder)+unsafe.Sizeof(inorder))
	bufio := new(bytes.Buffer)
	// 写入前序遍历结果
	err := binary.Write(bufio, binary.BigEndian, int32(len(preorder))) // 写入长度
	if err != nil {
		fmt.Println("binary.Write failed (preorder length):", err)
		return ""
	}
	for _, v := range preorder {
		if err := binary.Write(bufio, binary.BigEndian, int64(v)); err != nil {
			fmt.Println("binary.Write failed (preorder element):", err)
			return ""
		}
	}

	// 写入中序遍历结果
	err = binary.Write(bufio, binary.BigEndian, int32(len(inorder))) // 写入长度
	if err != nil {
		fmt.Println("binary.Write failed (inorder length):", err)
		return ""
	}
	for _, v := range inorder {
		if err := binary.Write(bufio, binary.BigEndian, int64(v)); err != nil {
			fmt.Println("binary.Write failed (inorder element):", err)
			return ""
		}
	}
	return string(bufio.Bytes())
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	bufio := bytes.NewBuffer([]byte(data))

	// 2. 解析前序遍历数组
	var preorderLen int32
	if err := binary.Read(bufio, binary.BigEndian, &preorderLen); err != nil {
		fmt.Println("binary.Read failed (preorder length):", err)
		return nil
	}

	preorder := make([]int64, preorderLen)
	for i := 0; i < int(preorderLen); i++ {
		if err := binary.Read(bufio, binary.BigEndian, &preorder[i]); err != nil {
			fmt.Println("binary.Read failed (preorder element):", err)
			return nil
		}
	}

	// 3. 解析中序遍历数组
	var inorderLen int32
	if err := binary.Read(bufio, binary.BigEndian, &inorderLen); err != nil {
		fmt.Println("binary.Read failed (inorder length):", err)
		return nil
	}

	inorder := make([]int64, inorderLen)
	for i := 0; i < int(inorderLen); i++ {
		if err := binary.Read(bufio, binary.BigEndian, &inorder[i]); err != nil {
			fmt.Println("binary.Read failed (inorder element):", err)
			return nil
		}
	}

	// 4. 根据遍历结果重建二叉树
	return RecoverBTPreInorderOrder64(preorder, inorder)
}

func bt_tree_preorder_recursion(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{node.Val}
	if node.Left != nil {
		rc = append(rc, bt_tree_preorder_recursion(node.Left)...)
	}
	if node.Right != nil {
		rc = append(rc, bt_tree_preorder_recursion(node.Right)...)
	}
	return rc
}

func bt_tree_inorder_recursion(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	rc := []int{}
	if node.Left != nil {
		rc = append(rc, bt_tree_inorder_recursion(node.Left)...)
	}
	rc = append(rc, node.Val)
	if node.Right != nil {
		rc = append(rc, bt_tree_inorder_recursion(node.Right)...)
	}
	return rc
}

func RecoverBTPreInorderOrder64(preorder []int64, inorder []int64) *TreeNode {
	if len(inorder) != len(preorder) || len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{}
	for i, v := range inorder {
		if v == preorder[0] {
			node.Val = int(v)
			if i > 0 {
				node.Left = RecoverBTPreInorderOrder64(preorder[1:i+1], inorder[:i])
			}
			if i < len(inorder)-1 {
				node.Right = RecoverBTPreInorderOrder64(preorder[i+1:], inorder[i+1:])
			}
			break
		}
	}
	return node
}

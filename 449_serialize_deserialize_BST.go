package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"
)

func TestSerializeDeserializeBST() {
	C := ConstructorCodecBST()
	fmt.Println(CompareTree(C.deserialize(C.serialize(nil)), nil))
	n00 := &TreeNode{Val: 10}
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n11 := &TreeNode{Val: 15}
	n00.Right = n11
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n23 := &TreeNode{Val: 20}
	n11.Right = n23
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n10 := &TreeNode{Val: 5}
	n00.Left = n10
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))
	n20 := &TreeNode{Val: 0}
	n10.Left = n20
	fmt.Println(CompareTree(C.deserialize(C.serialize(n00)), n00))

}

type CodecBST struct {
	b   *bytes.Buffer
	cur int
}

func ConstructorCodecBST() CodecBST {
	return CodecBST{}
}

// Serializes a tree to a single string.
func (this *CodecBST) serialize(root *TreeNode) string {
	this.b = &bytes.Buffer{}
	return this.serializeElement(root)
}

func (this *CodecBST) serializeElement(root *TreeNode) string {
	if root == nil {
		return ""
	}
	e := binary.Write(this.b, binary.BigEndian, int32(root.Val))
	if e != nil {
		fmt.Println(e.Error())
	}
	this.serializeElement(root.Left)
	this.serializeElement(root.Right)
	return this.b.String()
}

// Deserializes your encoded data to tree.
func (this *CodecBST) deserialize(data string) *TreeNode {
	d := make([]int32, len(data)/int(unsafe.Sizeof(int32(0))))
	this.b = bytes.NewBufferString(data)
	binary.Read(this.b, binary.BigEndian, &d)
	this.cur = 0
	return this.deserializeElement(d, math.MinInt32, math.MaxInt32)
}

func (this *CodecBST) deserializeElement(data []int32, lower, upper int32) *TreeNode {
	if this.cur >= len(data) {
		return nil
	}
	if data[this.cur] < lower || data[this.cur] > upper {
		return nil
	}
	tn := &TreeNode{Val: int(data[this.cur])}
	this.cur++
	tn.Left = this.deserializeElement(data, lower, int32(tn.Val))
	tn.Right = this.deserializeElement(data, int32(tn.Val), upper)
	return tn
}

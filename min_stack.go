package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

func TestMinStack() {
	s := NewHeapMinStack()
	fmt.Println(s.heap != nil)
	fmt.Println(s.tail.Len() == 0)

	s.Push(0)
	fmt.Println(s.tail.Len() == 1)
	fmt.Println(s.heap.Len() == 1)
	s.Pop()
	fmt.Println(s.tail.Len() == 0)
	fmt.Println(s.heap.Len() == 0)

	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top() == 2)
	fmt.Println(s.GetMin() == 1)

	s1 := NewHeapMinStack()
	s1.Push(-2)
	s1.Push(0)
	s1.Push(-3)
	fmt.Println(s1.GetMin() == -3)
	s1.Pop()
	fmt.Println(s1.Top() == 0)
	fmt.Println(s1.GetMin() == -2)
}

type MinHeapElement struct {
	val   int
	index int
}

type minHeap []*MinHeapElement

func (ms minHeap) Len() int {
	return len(ms)
}
func (ms minHeap) Less(i, j int) bool {
	return ms[i].val < ms[j].val
}
func (ms minHeap) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
	ms[i].index, ms[j].index = ms[j].index, ms[i].index
}
func (ms *minHeap) Push(val any) {
	*ms = append(*ms, val.(*MinHeapElement))
}
func (ms *minHeap) Pop() any {
	rc := (*ms)[ms.Len()-1]
	*ms = (*ms)[:ms.Len()-1]
	return rc
}

type MinStack struct {
	heap *minHeap
	tail list.List
}

func NewHeapMinStack() *MinStack {
	s := make(minHeap, 0)
	heap.Init(&s)
	return &MinStack{tail: list.List{}, heap: &s}
}

func (ms *MinStack) Push(val int) {
	v := &MinHeapElement{val: val, index: ms.heap.Len()}
	ms.tail.PushBack(v)
	heap.Push(ms.heap, v)
}

func (ms *MinStack) Pop() {
	rc := ms.tail.Back()
	if rc != nil {
		ms.tail.Remove(rc)
		heap.Remove(ms.heap, rc.Value.(*MinHeapElement).index)
	}
}

func (ms *MinStack) Top() int {
	rc := ms.tail.Back()
	//if rc == nil {
	//}
	return rc.Value.(*MinHeapElement).val
}

func (ms *MinStack) GetMin() int {
	//if ms.heap.Len() == 0 {
	//}
	return (*ms.heap)[0].val
}

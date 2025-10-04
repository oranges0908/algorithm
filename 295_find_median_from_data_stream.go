package main

import (
	"container/heap"
	"fmt"
)

func TestMedianFinder() {
	mf := ConstructorMedianFinder()
	fmt.Println(mf.FindMedian() == 0)
	mf.AddNum(1)
	fmt.Println(mf.FindMedian() == 1)
	mf.AddNum(-1)
	fmt.Println(mf.FindMedian() == 0)
	mf.AddNum(0)
	fmt.Println(mf.FindMedian() == 0)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian() == 0.5)
	mf.AddNum(3)
	fmt.Println(mf.FindMedian() == 1)
	mf.AddNum(-2)
	fmt.Println(mf.FindMedian() == 0.5)
	mf.AddNum(-5)
	fmt.Println(mf.FindMedian() == 0)
}

type intMinStack []int

func (s intMinStack) Len() int {
	return len(s)
}
func (s intMinStack) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s intMinStack) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *intMinStack) Push(x any) {
	*s = append(*s, x.(int))
}
func (s *intMinStack) Pop() any {
	rc := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return rc
}

type intMaxStack []int

func (s intMaxStack) Len() int {
	return len(s)
}
func (s intMaxStack) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s intMaxStack) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *intMaxStack) Push(x any) {
	*s = append(*s, x.(int))
}
func (s *intMaxStack) Pop() any {
	rc := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return rc
}

type MedianFinder struct {
	smallers *intMaxStack
	biggers  *intMinStack
}

func ConstructorMedianFinder() MedianFinder {
	mf := MedianFinder{}
	smallers := make(intMaxStack, 0)
	mf.smallers = &smallers
	heap.Init(mf.smallers)
	biggers := make(intMinStack, 0)
	mf.biggers = &biggers
	heap.Init(mf.biggers)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	if this.FindMedian() > float64(num) {
		heap.Push(this.smallers, num)
	} else {
		heap.Push(this.biggers, num)
	}
	this.balance()
}

func (this *MedianFinder) balance() {
	for len(*this.smallers)-len(*this.biggers) > 1 {
		heap.Push(this.biggers, heap.Pop(this.smallers))
	}
	for len(*this.biggers)-len(*this.smallers) > 1 {
		heap.Push(this.smallers, heap.Pop(this.biggers))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	s := 0
	if len(*this.smallers) > 0 {
		s = (*this.smallers)[0]
		if len(*this.smallers) > len(*this.biggers) {
			return float64(s)
		}
	}
	b := 0
	if len(*this.biggers) > 0 {
		b = (*this.biggers)[0]
		if len(*this.biggers) > len(*this.smallers) {
			return float64(b)
		}
	}
	return float64(s)/2 + float64(b)/2
}

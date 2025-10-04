package main

import "fmt"

func TestMyQueue() {
	mq := ConstructorMyQueue()
	fmt.Println(mq.Empty() == true)
	mq.Push(1)
	fmt.Println(mq.Empty() == false)
	fmt.Println(mq.Peek() == 1)
	mq.Push(2)
	fmt.Println(mq.Pop() == 1)
	fmt.Println(mq.Peek() == 2)
}

type MyQueue struct {
	inStack  MyStack
	outStack MyStack
}

func ConstructorMyQueue() MyQueue {
	mq := MyQueue{}
	mq.inStack = ConstructorMyStack()
	mq.outStack = ConstructorMyStack()
	return mq
}

func (this *MyQueue) Push(x int) {
	this.inStack.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.outStack.Empty() {
		for !this.inStack.Empty() {
			this.outStack.Push(this.inStack.Pop())
		}
	}
	return this.outStack.Pop()
}

func (this *MyQueue) Peek() int {
	if this.outStack.Empty() {
		for !this.inStack.Empty() {
			this.outStack.Push(this.inStack.Pop())
		}
	}
	return this.outStack.Top()
}

func (this *MyQueue) Empty() bool {
	return this.inStack.Empty() && this.outStack.Empty()
}

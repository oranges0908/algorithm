package main

import (
	"container/list"
	"fmt"
)

func TestMyStack() {
	ms := ConstructorMyStack()

	fmt.Println(ms.Empty() == true)

	ms.Push(1)
	fmt.Println(ms.Empty() == false)
	fmt.Println(ms.Top() == 1)

	ms.Push(2)
	fmt.Println(ms.Pop() == 2)
	fmt.Println(ms.Top() == 1)
}

type MyStack struct {
	data list.List
}

func ConstructorMyStack() MyStack {
	return MyStack{data: list.List{}}
}

func (this *MyStack) Push(x int) {
	this.data.PushBack(x)
}

func (this *MyStack) Pop() int {
	d := this.data.Back()
	this.data.Remove(d)
	return d.Value.(int)
}

func (this *MyStack) Top() int {
	d := this.data.Back()
	return d.Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.data.Len() == 0
}

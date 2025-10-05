package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func TestDailyTemperatures() {
	fmt.Println(reflect.DeepEqual(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0}))
	fmt.Println(reflect.DeepEqual(dailyTemperatures([]int{30, 40, 50, 60}), []int{1, 1, 1, 0}))
	fmt.Println(reflect.DeepEqual(dailyTemperatures([]int{30, 60, 90}), []int{1, 1, 0}))
	fmt.Println(reflect.DeepEqual(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}), []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}))
}

type dailyTemperaturesElem struct {
	id          int
	temperature int
}

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))

	monotonicStack := list.List{}
	for i := len(T) - 1; i >= 0; i-- {
		for monotonicStack.Len() > 0 && monotonicStack.Back().Value.(*dailyTemperaturesElem).temperature <= T[i] {
			monotonicStack.Remove(monotonicStack.Back())
		}

		if monotonicStack.Len() > 0 {
			result[i] = monotonicStack.Back().Value.(*dailyTemperaturesElem).id - i
		} else {
			result[i] = 0
		}

		monotonicStack.PushBack(&dailyTemperaturesElem{i, T[i]})
	}
	return result
}

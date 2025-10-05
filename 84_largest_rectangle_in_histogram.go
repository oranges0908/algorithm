package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

func TestLargestRectangleInHistogram() {
	fmt.Println(largestRectangleInHistogram([]int{}) == 0)
	fmt.Println(largestRectangleInHistogram([]int{2}) == 2)
	fmt.Println(largestRectangleInHistogram([]int{2, 2}) == 4)
	fmt.Println(largestRectangleInHistogram([]int{2, 4}) == 4)
	fmt.Println(largestRectangleInHistogram([]int{4, 1}) == 4)
	fmt.Println(largestRectangleInHistogram([]int{2, 1, 5, 6, 2, 3}) == 10)
}

func largestRectangleInHistogram(heights []int) int {
	return largestRectangleInHistogram3(heights)
	//return largestRectangleInHistogram1(heights, 0)
}

func largestRectangleInHistogram1(heights []int, pruneValue int) int {
	// Loop termination condition
	if len(heights) <= 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0] * 1
	}

	// find the minimum value
	a := math.MaxInt
	b := 0
	p := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] < a {
			a = heights[i]
			p = i
		}
		if heights[i] > b {
			b = heights[i]
		}
	}
	// calculate area
	m := len(heights) * a

	var l, r int
	if a != b && b*len(heights) > pruneValue {
		l = largestRectangleInHistogram1(heights[:p], max(m, pruneValue))
		r = largestRectangleInHistogram1(heights[p+1:], max(m, pruneValue))
	}
	//
	return max(m, l, r)
}

func largestRectangleInHistogram2(heights []int) int {
	// generate minimum heap
	mh := minimumHeap{}
	heap.Init(&mh)
	for i, h := range heights {
		heap.Push(&mh, heapElement{h, i})
	}

	// generate cache
	// BST is best
	m := make(map[int]int)
	m[0] = len(heights) - 1
	area := 0

	// main loop
	for mh.Len() > 0 {
		min := heap.Pop(&mh).(heapElement)
		index := min.data.(int)
		for s, e := range m {
			if s <= index && index <= e {
				//calculate area
				a := (e - s + 1) * min.value
				if a > area {
					area = a
				}

				// update cache
				delete(m, s)
				if index > s {
					m[s] = index - 1
				}
				if index < e {
					m[index+1] = e
				}
				break
			}
		}
	}
	return area
}

type positionValue struct {
	value    int
	position int
}

func largestRectangleInHistogram3(heights []int) int {
	l := list.List{}
	leftBoundry := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		for l.Len() > 0 && l.Back().Value.(positionValue).value >= heights[i] {
			l.Remove(l.Back())
		}

		if l.Len() > 0 {
			leftBoundry[i] = l.Back().Value.(positionValue).position
		} else {
			leftBoundry[i] = -1
		}

		l.PushBack(positionValue{value: heights[i], position: i})
	}

	l = list.List{}
	rightBoundry := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for l.Len() > 0 && l.Back().Value.(positionValue).value >= heights[i] {
			l.Remove(l.Back())
		}

		if l.Len() > 0 {
			rightBoundry[i] = l.Back().Value.(positionValue).position
		} else {
			rightBoundry[i] = len(heights)
		}

		l.PushBack(positionValue{value: heights[i], position: i})
	}

	area := 0
	for i := 0; i < len(heights); i++ {
		a := (rightBoundry[i] - leftBoundry[i] - 1) * heights[i]
		if a > area {
			area = a
		}
	}
	return area
}

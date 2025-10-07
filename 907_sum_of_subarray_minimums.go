package main

import (
	"container/list"
	"fmt"
)

func TestSumSubarrayMins() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}) == 17)
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}) == 444)
	fmt.Println(sumSubarrayMins([]int{71, 55, 82, 55}) == 593)
}

type element2 struct {
	value int
	index int
}

func sumSubarrayMins(A []int) int {
	leftLess := make([]int, len(A))
	l := list.New()
	for j := 0; j < len(A); j++ {
		for l.Len() > 0 && l.Back().Value.(element2).value > A[j] {
			l.Remove(l.Back())
		}
		if l.Len() > 0 {
			leftLess[j] = l.Back().Value.(element2).index
		} else {
			leftLess[j] = -1
		}
		l.PushBack(element2{A[j], j})
	}

	rightLess := make([]int, len(A))
	l = list.New()

	for j := len(A) - 1; j >= 0; j-- {
		for l.Len() > 0 && l.Back().Value.(element2).value >= A[j] {
			l.Remove(l.Back())
		}
		if l.Len() > 0 {
			rightLess[j] = l.Back().Value.(element2).index
		} else {
			rightLess[j] = len(A)
		}
		l.PushBack(element2{A[j], j})
	}

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += (rightLess[i] - i) * (i - leftLess[i]) * A[i]
		sum = sum % (1000000000 + 7)
	}
	return sum
}

// time complexity O(n^2)
func sumSubarrayMins1(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		l := list.New()
		for j := i; j < len(A); j++ {
			for l.Len() > 0 && l.Back().Value.(int) > A[j] {
				l.Remove(l.Back())
			}
			l.PushBack(A[j])
			sum += l.Front().Value.(int)
			sum = sum % (1000000000 + 7)
		}
	}
	return sum
}

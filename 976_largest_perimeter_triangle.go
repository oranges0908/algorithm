package main

import (
	"fmt"
	"sort"
)

func TestLargestPerimeterTriangle() {
	a := []int{2, 1, 2}
	fmt.Println(largestPerimeterTriangle(a) == 5)

	a = []int{1, 2, 1, 10}
	fmt.Println(largestPerimeterTriangle(a) == 0)

	a = []int{1, 2, 9, 10}
	fmt.Println(largestPerimeterTriangle(a) == 21)

}

type sortInt []int

func (s sortInt) Len() int {
	return len(s)
}

func (s sortInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestPerimeterTriangle(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Sort(sortInt(A))
	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

package main

import (
	"fmt"
	"slices"
)

func testSort() {
	tf := QuickSort

	arr1 := []int{5, 2, 3, 1}
	arr1t := []int{1, 2, 3, 5}
	fmt.Println(slices.Equal(tf(arr1), arr1t))

	arr2 := []int{5, 1, 1, 2, 0, 0}
	arr2t := []int{0, 0, 1, 1, 2, 5}
	fmt.Println(slices.Equal(tf(arr2), arr2t))
}

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		maxValueIndex := i
		for j := 0; j < i; j++ {
			if arr[j] > arr[maxValueIndex] {
				maxValueIndex = j
			}
		}
		arr[maxValueIndex], arr[i] = arr[i], arr[maxValueIndex]
	}
	return arr
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	rc1 := make([]int, 0)
	rc2 := make([]int, 0)
	for j := 1; j < len(arr); j++ {
		if arr[j] < arr[0] {
			rc1 = append(rc1, arr[j])
		} else {
			rc2 = append(rc2, arr[j])
		}
	}

	return append(append(QuickSort(rc1), arr[0]), QuickSort(rc2)...)
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	rc := make([]int, 0, len(arr))
	rc1 := MergeSort(arr[0 : len(arr)/2])
	rc2 := MergeSort(arr[len(arr)/2:])
	i := 0
	j := 0
	for i < len(rc1) && j < len(rc2) {
		if rc1[i] <= rc2[j] {
			rc = append(rc, rc1[i])
			i++
		} else {
			rc = append(rc, rc2[j])
			j++
		}
	}
	if i < len(rc1) {
		rc = append(rc, rc1[i:]...)
	}
	if j < len(rc2) {
		rc = append(rc, rc2[j:]...)
	}

	return rc
}

func HeapSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		maxValueIndex := i
		for j := 0; j < i; j++ {
			if arr[j] > arr[maxValueIndex] {
				maxValueIndex = j
			}
		}
		arr[maxValueIndex], arr[i] = arr[i], arr[maxValueIndex]
	}
	return arr
}

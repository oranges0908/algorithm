package main

import "fmt"

func testFindMedianSortedArrays() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{}) == 0.0)
	fmt.Println(findMedianSortedArrays([]int{1}, []int{}) == 1.0)
	fmt.Println(findMedianSortedArrays([]int{}, []int{2}) == 2.0)
	fmt.Println(findMedianSortedArrays([]int{0}, []int{2}) == 1.0)
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}) == 2.0)
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}) == 2.5)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}

	mergeArray := []int{}
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			mergeArray = append(mergeArray, nums1[i])
			i++
		} else {
			mergeArray = append(mergeArray, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		mergeArray = append(mergeArray, nums1[i:]...)
	}
	if j < len(nums2) {
		mergeArray = append(mergeArray, nums2[j:]...)
	}

	total := len(mergeArray)
	if total%2 == 0 {
		return float64(mergeArray[total/2]+mergeArray[total/2-1]) / 2
	} else {
		return float64(mergeArray[total/2])
	}
}

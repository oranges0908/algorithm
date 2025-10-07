package main

import (
	"container/list"
	"fmt"
)

func TestTrappingRainWater() {
	fmt.Println(trappingRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) == 6)
	fmt.Println(trappingRainWater([]int{4, 2, 0, 3, 2, 5}) == 9)
	fmt.Println(trappingRainWater([]int{1, 1, 1, 1}) == 0)
	fmt.Println(trappingRainWater([]int{1, 2, 1, 1}) == 0)
	fmt.Println(trappingRainWater([]int{2, 0, 2}) == 2)
}

func trappingRainWater(heightMap []int) int {
	rc := 0
	l := list.New()
	for i := 0; i < len(heightMap); i++ {
		for l.Len() > 0 && heightMap[l.Back().Value.(int)] <= heightMap[i] {
			v := l.Back().Value.(int)
			l.Remove(l.Back())
			if l.Len() > 0 {
				p := l.Back().Value.(int)
				rc += (min(heightMap[i], heightMap[p]) - heightMap[v]) * (i - p - 1)
			}
		}
		l.PushBack(i)
	}
	return rc
}

func trappingRainWater2(heightMap []int) int {
	rc := 0

	maxLeft := 0
	maxRight := 0

	i := 0
	j := len(heightMap) - 1
	for i < j {
		if heightMap[i] < heightMap[j] {
			if heightMap[i] > maxLeft {
				maxLeft = heightMap[i]
			} else {
				rc += min(maxLeft, heightMap[j]) - heightMap[i]
			}
			i++
		} else {
			if heightMap[j] > maxRight {
				maxRight = heightMap[j]
			} else {
				rc += min(maxRight, heightMap[i]) - heightMap[j]
			}
			j--
		}
	}
	return rc
}

func trappingRainWater1(heightMap []int) int {
	//求区域最高
	peeks := make([]int, 0)
	l := list.New()
	for i := 0; i < len(heightMap); i++ {
		for l.Len() > 0 && heightMap[l.Back().Value.(int)] <= heightMap[i] {
			l.Remove(l.Back())
		}
		l.PushBack(i)
		if l.Len() == 1 {
			peeks = append(peeks, l.Back().Value.(int))
		}
	}

	rc := 0
	for i := 0; i+1 < len(peeks); i++ {
		r := min(heightMap[peeks[i]], heightMap[peeks[i+1]]) * (peeks[i+1] - peeks[i] - 1)
		for j := peeks[i] + 1; j < peeks[i+1]; j++ {
			r -= heightMap[j]
		}
		rc += r
	}

	start := peeks[len(peeks)-1]
	peeks = make([]int, 0)
	l = list.New()
	for i := len(heightMap) - 1; i >= start; i-- {
		for l.Len() > 0 && heightMap[l.Back().Value.(int)] <= heightMap[i] {
			l.Remove(l.Back())
		}
		l.PushBack(i)
		if l.Len() == 1 {
			peeks = append(peeks, l.Back().Value.(int))
		}
	}

	for i := 0; i+1 < len(peeks); i++ {
		r := min(heightMap[peeks[i]], heightMap[peeks[i+1]]) * (peeks[i] - peeks[i+1] - 1)
		for j := peeks[i+1] + 1; j < peeks[i]; j++ {
			r -= heightMap[j]
		}
		rc += r
	}
	return rc
}

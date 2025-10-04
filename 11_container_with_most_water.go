package main

import "fmt"

func TestMaxArea() {
	fmt.Println(maxArea([]int{1, 1}) == 1)
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
}

func maxArea1(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max_area := 0
	last_height_i := 0
	for i := 0; i < len(height); i++ {
		// update i
		if height[i] <= last_height_i {
			continue
		}
		last_height_i = height[i]

		last_high_j := 0
		for j := len(height) - 1; j > i; j-- {
			// only calculate bigger height[j]
			if last_high_j < height[j] {
				ht := height[i]
				if height[j] < ht {
					ht = height[j]
				}

				area := ht * (j - i)
				if area > max_area {
					max_area = area
				}

				// update last_high_j
				last_high_j = height[j]

				if height[j] > height[i] {
					break
				}
			}
		}
	}
	return max_area
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max_area := 0
	i := 0
	j := len(height) - 1
	for i < j {
		ht := height[i]
		if height[j] < ht {
			ht = height[j]
		}

		area := ht * (j - i)
		if area > max_area {
			max_area = area
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max_area
}

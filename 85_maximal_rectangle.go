package main

import "fmt"

func TestMaximalRectangle() {
	m := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(m) == 6)

	m = [][]byte{{'0'}}
	fmt.Println(maximalRectangle(m) == 0)

	m = [][]byte{{'1'}}
	fmt.Println(maximalRectangle(m) == 1)
}

func maximalRectangle(matrix [][]byte) int {
	//结合84题，将其转换为多组柱状图，分别计算
	mtx := make([][]int, 0)
	r1 := make([]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		r1 = append(r1, int(matrix[0][i])-'0')
	}
	mtx = append(mtx, r1)
	for i := 1; i < len(matrix); i++ {
		r := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				r[j] = mtx[i-1][j] + 1
			} else {
				r[j] = 0
			}
		}
		mtx = append(mtx, r)
	}

	//fmt.Println("calculate result")
	//printMatrix(mtx)

	m := 0
	for i := 0; i < len(mtx); i++ {
		t := largestRectangleInHistogram3(mtx[i])
		if t > m {
			m = t
		}
	}

	return m
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

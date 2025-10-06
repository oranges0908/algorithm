package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"reflect"
)

func TestPacificAtlanticWaterFlow() {
	heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	Output := [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}
	checkPacificAtlanticWaterFlowResult(heights, Output)

	heights = [][]int{{1}}
	Output = [][]int{{0, 0}}
	checkPacificAtlanticWaterFlowResult(heights, Output)

	heights = [][]int{{1, 2}, {4, 3}}
	Output = [][]int{{0, 1}, {1, 0}, {1, 1}}
	checkPacificAtlanticWaterFlowResult(heights, Output)

	heights = [][]int{{1, 1}, {1, 1}}
	Output = [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	checkPacificAtlanticWaterFlowResult(heights, Output)
}

func checkPacificAtlanticWaterFlowResult(heights [][]int, Output [][]int) {
	rc := pacificAtlanticWaterFlow(heights)
	fmt.Println(Output)
	fmt.Println(rc)
outloop:
	for i := range Output {
		for j := range rc {
			if reflect.DeepEqual(Output[i], rc[j]) {
				fmt.Println(Output[i], "check ok")
				continue outloop
			}
		}
		fmt.Println(Output[i], "check failed")
	}
}

func pacificAtlanticWaterFlow(flow [][]int) [][]int {
	return pacificAtlanticWaterFlow_deepSearch(flow)
}

type Position struct {
	x int
	y int
}

func pacificAtlanticWaterFlow_deepSearch(flow [][]int) [][]int {
	pacificCache := make(map[Position]struct{})
	atlanticCache := make(map[Position]struct{})

	todo := list.List{}
	for i := 0; i < len(flow); i++ {
		todo.PushBack(Position{i, 0})
	}
	for i := 0; i < len(flow[0]); i++ {
		todo.PushBack(Position{0, i})
	}
	connectionDetector(flow, &todo, pacificCache)

	todo = list.List{}
	for i := 0; i < len(flow); i++ {
		todo.PushBack(Position{i, len(flow[0]) - 1})
	}
	for i := 0; i < len(flow[0]); i++ {
		todo.PushBack(Position{len(flow) - 1, i})
	}
	connectionDetector(flow, &todo, atlanticCache)

	rc := make([][]int, 0)
	for i := 0; i < len(flow); i++ {
		for j := 0; j < len(flow[0]); j++ {
			_, ok1 := pacificCache[Position{i, j}]
			_, ok2 := atlanticCache[Position{i, j}]
			if ok1 && ok2 {
				rc = append(rc, []int{i, j})
			}
		}
	}
	return rc
}

func connectionDetector(flow [][]int, todo *list.List, rc map[Position]struct{}) {
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for todo.Len() > 0 {
		a := todo.Front()
		p := a.Value.(Position)
		todo.Remove(a)

		rc[p] = struct{}{}

		for _, direction := range directions {
			x := p.x + direction[0]
			y := p.y + direction[1]
			if x < 0 || x >= len(flow) || y < 0 || y >= len(flow[0]) {
				continue
			}
			if _, ok := rc[Position{x, y}]; ok {
				continue
			}

			if flow[p.x][p.y] <= flow[x][y] {
				todo.PushBack(Position{x, y})
			}
		}
	}
}

func pacificAtlanticWaterFlow_monotonic(flow [][]int) [][]int {
	mh := make(minimumHeap, 0)
	for x, r := range flow {
		for y, v := range r {
			heap.Push(&mh, heapElement{value: v, data: Position{x, y}})
		}
	}

	rc := make(map[[2]int][2]bool, 0)
	for mh.Len() > 0 {
		p := heap.Pop(&mh).(heapElement).data.(Position)
		checkOceanConnect(flow, p.x, p.y, rc)
	}

	r := make([][]int, 0)
	for k, v := range rc {
		if v[0] && v[1] {
			r = append(r, []int{k[0], k[1]})
		}
	}
	return r
}

func checkOceanConnect(flow [][]int, x, y int, rc map[[2]int][2]bool) {
	var up, left, down, right [2]bool

	if x == 0 {
		up = [2]bool{true, false}
	} else {
		if v, ok := rc[[2]int{x - 1, y}]; ok {
			up = v
		}
	}

	if y == 0 {
		left = [2]bool{true, false}
	} else {
		if v, ok := rc[[2]int{x, y - 1}]; ok {
			left = v
		}
	}

	if x == len(flow)-1 {
		down = [2]bool{false, true}
	} else {
		if v, ok := rc[[2]int{x + 1, y}]; ok {
			down = v
		}
	}

	if y == len(flow[0])-1 {
		right = [2]bool{false, true}
	} else {
		if v, ok := rc[[2]int{x, y + 1}]; ok {
			right = v
		}
	}

	r := [2]bool{false, false}
	r[0] = up[0] || left[0] || down[0] || right[0]
	r[1] = up[1] || left[1] || down[1] || right[1]
	rc[[2]int{x, y}] = r
}

func pacificAtlanticWaterFlow_deepSearchDP(flow [][]int) [][]int {
	cache := make(map[[2]int][2]int)
	paths := make(map[[2]int]struct{}, 0)
	for i := 0; i < len(flow); i++ {
		for j := len(flow[i]) - 1; j > 0; j-- {
			checkWaterFlow(flow, [2]int{i, j}, cache, paths)
		}
	}

	rc := make([][]int, 0)
	for i := 0; i < len(flow); i++ {
		for j := len(flow[i]) - 1; j > 0; j-- {
			if cache[[2]int{i, j}][0] == 1 && cache[[2]int{i, j}][1] == 1 {
				rc = append(rc, []int{i, j})
			}
		}
	}
	return rc
}

func checkWaterFlow(flow [][]int, point [2]int, cache map[[2]int][2]int, paths map[[2]int]struct{}) [2]int {
	if v, ok := cache[point]; ok {
		return v
	}

	var up, left, down, right [2]int
	upPoint := [2]int{point[0] - 1, point[1]}
	if point[0] == 0 {
		up = [2]int{1, -1}
	} else {
		if flow[upPoint[0]][upPoint[1]] <= flow[point[0]][point[1]] {
			if _, ok := paths[upPoint]; !ok {
				paths[upPoint] = struct{}{}
				up = checkWaterFlow(flow, upPoint, cache, paths)
				delete(paths, upPoint)
			}
		} else {
			up = [2]int{-1, -1}
		}
	}

	leftPoint := [2]int{point[0], point[1] - 1}
	if point[1] == 0 {
		left = [2]int{1, -1}
	} else {
		if flow[leftPoint[0]][leftPoint[1]] <= flow[point[0]][point[1]] {
			if _, ok := paths[leftPoint]; !ok {
				paths[leftPoint] = struct{}{}
				left = checkWaterFlow(flow, leftPoint, cache, paths)
				delete(paths, leftPoint)
			}
		} else {
			left = [2]int{-1, -1}
		}
	}

	downPoint := [2]int{point[0] + 1, point[1]}
	if point[0] == 4 {
		down = [2]int{-1, 1}
	} else {
		if flow[downPoint[0]][downPoint[1]] <= flow[point[0]][point[1]] {
			if _, ok := paths[downPoint]; !ok {
				paths[downPoint] = struct{}{}
				down = checkWaterFlow(flow, downPoint, cache, paths)
				delete(paths, downPoint)
			}
		} else {
			down = [2]int{-1, -1}
		}
	}

	rightPoint := [2]int{point[0], point[1] + 1}
	if point[1] == 4 {
		right = [2]int{-1, 1}
	} else {
		if flow[rightPoint[0]][rightPoint[1]] <= flow[point[0]][point[1]] {
			if _, ok := paths[rightPoint]; !ok {
				paths[rightPoint] = struct{}{}
				right = checkWaterFlow(flow, rightPoint, cache, paths)
				delete(paths, rightPoint)
			}
		} else {
			right = [2]int{-1, -1}
		}
	}

	rc := [2]int{}
	if up[0] == 1 || left[0] == 1 || down[0] == 1 || right[0] == 1 {
		rc[0] = 1
	} else if up[0]+left[0]+down[0]+right[0] == -4 {
		rc[0] = -1
	}
	if up[1] == 1 || left[1] == 1 || down[1] == 1 || right[1] == 1 {
		rc[1] = 1
	} else if up[1]+left[1]+down[1]+right[1] == -4 {
		rc[1] = -1
	}

	if rc[0]*rc[1] != 0 {
		cache[point] = rc
	}
	return rc
}

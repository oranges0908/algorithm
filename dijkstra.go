package main

import (
	"container/heap"
	"fmt"
)

func TestDijkstra() {
	graph := [][]int{{0, 1, 2},
		{0, 2, 6},
		{1, 3, 5},
		{2, 3, 8},
		{3, 4, 10},
		{3, 5, 15},
		{4, 5, 6},
		{4, 6, 2},
		{6, 5, 6}}

	graph = appendReturnRoad(graph)
	fmt.Println(DijkstraHeap(7, graph, 0, 0) == 0)
	fmt.Println(DijkstraHeap(7, graph, 0, 1) == 2)
	fmt.Println(DijkstraHeap(7, graph, 0, 2) == 6)
	fmt.Println(DijkstraHeap(7, graph, 0, 3) == 7)
	fmt.Println(DijkstraHeap(7, graph, 0, 4) == 17)
	fmt.Println(DijkstraHeap(7, graph, 0, 5) == 22)
	fmt.Println(DijkstraHeap(7, graph, 0, 6) == 19)

	fmt.Println("DeepDijkstraHeap")
	graph = [][]int{{0, 1, 100},
		{1, 2, 100},
		{1, 3, 600},
		{2, 3, 200},
		{2, 0, 100}}
	fmt.Println(DeepDijkstraHeap(3, graph, 0, 1, 2) == 100)
	fmt.Println(DeepDijkstraHeap(3, graph, 0, 2, 2) == 200)
	fmt.Println(DeepDijkstraHeap(3, graph, 0, 3, 2) == 400)

	fmt.Println(DeepDijkstraHeap(3, graph, 0, 1, 1) == 100)
	fmt.Println(DeepDijkstraHeap(3, graph, 0, 2, 0) == -1)
	fmt.Println(DeepDijkstraHeap(3, graph, 0, 3, 1) == 700)

	fmt.Println("DeepDijkstraHeap")
	graph = [][]int{{3, 4, 4}, {2, 5, 6}, {4, 7, 10}, {9, 6, 5}, {7, 4, 4}, {6, 2, 10},
		{6, 8, 6}, {7, 9, 4}, {1, 5, 4}, {1, 0, 4}, {9, 7, 3}, {7, 0, 5}, {6, 5, 8},
		{1, 7, 6}, {4, 0, 9}, {5, 9, 1}, {8, 7, 3}, {1, 2, 6}, {4, 1, 5}, {5, 2, 4},
		{1, 9, 1}, {7, 8, 10}, {0, 4, 2}, {7, 2, 8}}

	fmt.Println(DeepDijkstraHeap(10, graph, 6, 5, 7) == 8)
	fmt.Println(DeepDijkstraHeap(10, graph, 6, 9, 7) == 9)
	fmt.Println(DeepDijkstraHeap(10, graph, 6, 0, 7) == 14)

	fmt.Println("DeepDijkstraHeap")
	graph = [][]int{{1, 0, 5}, {2, 1, 5}, {3, 0, 2}, {1, 3, 2}, {4, 1, 1}, {2, 4, 1}}
	fmt.Println(DeepDijkstraHeap(5, graph, 2, 0, 2) == 7)

	fmt.Println("DeepDijkstraHeap")
	graph = [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}
	fmt.Println(DeepDijkstraHeap(5, graph, 0, 2, 2) == 7)
}

func appendReturnRoad(g [][]int) [][]int {
	ng := make([][]int, 0)
	for _, r := range g {
		ng = append(ng, r)
		ng = append(ng, []int{r[1], r[0], r[2]})
	}
	return ng
}

type heapElement struct {
	value int
	data  any
}
type minimumHeap []heapElement

func (h minimumHeap) Len() int { return len(h) }
func (h minimumHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h minimumHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minimumHeap) Push(x any) {
	*h = append(*h, x.(heapElement))
}
func (h *minimumHeap) Pop() any {
	rc := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return rc
}

func DijkstraHeap(n int, graph [][]int, src, dst int) int {
	// convert edge list to adjacency map
	fullConnection := make(map[int]map[int]int)
	for _, w := range graph {
		if _, ok := fullConnection[w[0]]; !ok {
			fullConnection[w[0]] = make(map[int]int)
		}
		fullConnection[w[0]][w[1]] = w[2]
	}

	// min-heap to select the next node with the smallest distance
	toVisit := make(minimumHeap, 0)
	heap.Init(&toVisit)
	heap.Push(&toVisit, heapElement{value: 0, data: src})

	// record visited nodes for pruning redundant paths
	visitedNode := make(map[int]struct{})

	for toVisit.Len() > 0 {
		// get next node to visit
		item := heap.Pop(&toVisit).(heapElement)
		node := item.data.(int)

		// find destination
		if node == dst {
			return item.value
		}

		//mark the node as visited
		if _, ok := visitedNode[node]; ok {
			continue
		}
		visitedNode[node] = struct{}{}

		// check the neighbors of node, update the node need to visit
		for n, d := range fullConnection[node] {
			heap.Push(&toVisit, heapElement{value: item.value + d, data: n})
		}
	}
	return -1
}

type deepPath struct {
	dst   int
	depth int
}

func DeepDijkstraHeap(n int, graph [][]int, src int, dst int, depth int) int {
	// convert edge list to adjacency map
	fullConnection := make(map[int]map[int]int)
	for _, w := range graph {
		if _, ok := fullConnection[w[0]]; !ok {
			fullConnection[w[0]] = make(map[int]int)
		}
		fullConnection[w[0]][w[1]] = w[2]
	}

	// min-heap to select the next node with the smallest distance
	toVisit := make(minimumHeap, 0)
	heap.Init(&toVisit)
	heap.Push(&toVisit, heapElement{value: 0, data: deepPath{dst: src, depth: 0}})

	// record the node and path depth for pruning
	pruningMap := make(map[int]int)
	for toVisit.Len() > 0 {
		// get next node to visit
		item := heap.Pop(&toVisit).(heapElement)
		node := item.data.(deepPath).dst
		dp := item.data.(deepPath).depth

		// find the destination
		if node == dst {
			return item.value
		}

		// path depth check
		if dp > depth {
			continue
		}

		//Since we always process the smallest cost first with the min-heap,
		//if the current node comes with a higher cost and deeper path, it’s invalid and should be discarded.
		if v, ok := pruningMap[node]; ok && v <= dp {
			continue
		}
		pruningMap[node] = dp

		// check the neighbors of node, update the node need to visit
		for n, d := range fullConnection[node] {
			heap.Push(&toVisit, heapElement{value: item.value + d, data: deepPath{dst: n, depth: dp + 1}})
		}
	}
	return -1
}

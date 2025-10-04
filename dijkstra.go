package main

import (
	"container/heap"
	"fmt"
	"math"
)

func TestDijkstra() {
	graph := []WeightEdge{{0, 1, 2},
		{0, 2, 6},
		{1, 3, 5},
		{2, 3, 8},
		{3, 4, 10},
		{3, 5, 15},
		{4, 5, 6},
		{4, 6, 2},
		{6, 5, 6}}

	graph = appendReturnRoad(graph)

	fmt.Println(graph)
	rc := Dijkstra(graph, 0)
	fmt.Println(rc)
	fmt.Println(rc[0] == 0)
	fmt.Println(rc[1] == 2)
	fmt.Println(rc[2] == 6)
	fmt.Println(rc[3] == 7)
	fmt.Println(rc[4] == 17)
	fmt.Println(rc[5] == 22)
	fmt.Println(rc[6] == 19)
}

func appendReturnRoad(g []WeightEdge) []WeightEdge {
	ng := make([]WeightEdge, 0)
	for _, r := range g {
		ng = append(ng, r)
		ng = append(ng, WeightEdge{r.Dst, r.Src, r.Weight})
	}
	return ng
}

type WeightEdge struct {
	Src    int
	Dst    int
	Weight int
}

func Dijkstra(graph []WeightEdge, src int) map[int]int {
	return DijkstraArray(graph, src)
}

func DijkstraArray(graph []WeightEdge, src int) map[int]int {
	minDistance := make(map[int]struct{})
	visitedNode := make(map[int]int)

	// transfer group to map
	fullConnection := make(map[int]map[int]int)
	for _, w := range graph {
		if _, ok := fullConnection[w.Src]; !ok {
			fullConnection[w.Src] = make(map[int]int)
		}
		fullConnection[w.Src][w.Dst] = w.Weight
	}

	// cache node to src distance
	visitedNode[src] = 0
	// cache minimum distance info
	minDistance[src] = struct{}{}

	for {
		// iterate the neighbor of minimum distance group
		for s, _ := range minDistance {
			for d, w := range fullConnection[s] {
				//skip member of minimum distance group
				if _, ok := minDistance[d]; ok {
					continue
				}
				// update distance info
				distance := visitedNode[s] + w
				if v, ok := visitedNode[d]; !ok || v > distance {
					visitedNode[d] = distance
				}
			}
		}
		// find the new smallest one
		newMinimum := src
		miniDistance := math.MaxInt
		for d, w := range visitedNode {
			if _, ok := minDistance[d]; ok {
				continue
			}
			if w < miniDistance {
				miniDistance = w
				newMinimum = d
			}
		}
		if newMinimum != src {
			minDistance[newMinimum] = struct{}{}
			continue
		}
		return visitedNode
	}
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

func DijkstraHeap(graph []WeightEdge, src int) map[int]int {
	nodes := make(map[int]struct{})

	// transfer group to map
	fullConnection := make(map[int]map[int]int)
	for _, w := range graph {
		if _, ok := fullConnection[w.Src]; !ok {
			fullConnection[w.Src] = make(map[int]int)
		}
		fullConnection[w.Src][w.Dst] = w.Weight
		nodes[w.Src] = struct{}{}
		nodes[w.Dst] = struct{}{}
	}

	visitedNode := make(map[int]struct{})
	nodeDistance := make(map[int]int)
	toVisit := make(minimumHeap, 0)
	heap.Init(&toVisit)
	heap.Push(&toVisit, heapElement{value: 0, data: src})

	for toVisit.Len() > 0 {
		item := heap.Pop(&toVisit).(heapElement)
		node := item.data.(int)

		if _, ok := visitedNode[node]; ok {
			continue
		}
		visitedNode[node] = struct{}{}

		for n, d := range fullConnection[node] {
			if nodeDistance[n] > nodeDistance[node]+d {
				nodeDistance[n] = nodeDistance[node] + d
				heap.Push(&toVisit, heapElement{value: nodeDistance[n], data: n})
			}
		}
	}
	return nodeDistance
}

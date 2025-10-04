package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	fm := newFlightMaps(flights)
	return fm.getCheapestPrice(src, dst, k)
}

func TestNewFlightMaps() {
	fp := newFlightMaps([][]int{{0, 1, 100}})
	fmt.Println("test 0")
	fmt.Println(fp.flights[0][1] == 100)
	fmt.Println(fp.getCheapestPrice(0, 2, 1) == -1)
	fmt.Println(fp.getCheapestPrice(0, 1, 0) == 100)
	fmt.Println(fp.getCheapestPrice(0, 1, 1) == 100)
	fmt.Println(fp.getCheapestPrice(0, 1, 2) == 100)
	fmt.Println(fp.getCheapestPrice(1, 0, 0) == -1)

	fp = newFlightMaps([][]int{{0, 1, 200}, {1, 0, 200}, {0, 1, 100}})
	fmt.Println("test 1")
	fmt.Println(fp.flights[1][0] == 200)
	fmt.Println(fp.getCheapestPrice(0, 1, 0) == 100)
	fmt.Println(fp.getCheapestPrice(1, 0, 0) == 200)

	fp = newFlightMaps([][]int{{0, 1, 100}, {1, 0, 200}, {1, 2, 200}})
	fmt.Println("test 2")
	fmt.Println(fp.flights[1][0] == 200)
	fmt.Println(fp.getCheapestPrice(0, 2, 0) == -1)
	fmt.Println(fp.getCheapestPrice(0, 2, 1) == 300)
	fmt.Println(fp.getCheapestPrice(0, 2, 2) == 300)
	fmt.Println(fp.getCheapestPrice(2, 0, 2) == -1)

	fp = newFlightMaps([][]int{{0, 1, 100}, {1, 2, 200}, {2, 3, 200}, {3, 0, 200}})
	fmt.Println("test 3")
	fmt.Println(fp.getCheapestPrice(0, 3, 1) == -1)
	fmt.Println(fp.getCheapestPrice(0, 3, 2) == 500)
	fmt.Println(fp.getCheapestPrice(1, 0, 1) == -1)
	fmt.Println(fp.getCheapestPrice(1, 0, 2) == 600)
	fmt.Println(fp.getCheapestPrice(1, 0, 3) == 600)

	fp = newFlightMaps([][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}})
	fmt.Println("test 4")
	fmt.Println(fp.getCheapestPrice(0, 3, 1) == 700)

	fp = newFlightMaps([][]int{{0, 1, 100}, {0, 2, 100}, {0, 3, 10}, {1, 2, 100}, {1, 4, 10}, {2, 1, 10}, {2, 3, 100}, {2, 4, 100}, {3, 2, 10}, {3, 4, 100}})
	fmt.Println("test 5")
	fmt.Println(fp.getCheapestPrice(0, 4, 3) == 40)

	fp = newFlightMaps([][]int{{1, 0, 5}, {2, 1, 5}, {3, 0, 2}, {1, 3, 2}, {4, 1, 1}, {2, 4, 1}})
	fmt.Println("test 6")
	fmt.Println(fp.getCheapestPrice(2, 0, 2) == 7)
}

type visitedNodeInfo struct {
	id         int
	totalPrice int
}

type flightMaps struct {
	flights         map[int]map[int]int
	cacheSearchPath []visitedNodeInfo
	miniTotalPrice  map[int]map[int]int
}

func newFlightMaps(flights [][]int) *flightMaps {
	fm := &flightMaps{}
	fm.flights = make(map[int]map[int]int)
	fm.cacheSearchPath = make([]visitedNodeInfo, 0)
	fm.miniTotalPrice = make(map[int]map[int]int)
	// initialize the flight map
	for _, routes := range flights {
		if len(routes) < 3 {
			continue
		}
		if _, ok := fm.flights[routes[0]]; !ok {
			fm.flights[routes[0]] = make(map[int]int)
		}
		// just save the lowest price
		if p, ok := fm.flights[routes[0]][routes[1]]; !ok || p > routes[2] {
			fm.flights[routes[0]][routes[1]] = routes[2]
		}
	}
	return fm
}

func (fm *flightMaps) getCheapestPrice(src int, dst int, k int) int {
	fm.miniTotalPrice = make(map[int]map[int]int)
	fm.calculateCheapestPrice(src, dst, k)
	if v, ok := fm.miniTotalPrice[dst]; ok {
		min := math.MaxInt
		for _, p := range v {
			if p < min {
				min = p
			}
		}
		return min
	}
	return -1
}

func (fm *flightMaps) calculateCheapestPrice(src int, dst int, k int) {

	// Depth-First Search to find the minimum cost
depthLoop:
	for ct, price := range fm.flights[src] {
		totalPrice := price
		if len(fm.cacheSearchPath) > 0 {
			totalPrice += fm.cacheSearchPath[len(fm.cacheSearchPath)-1].totalPrice
		}
		// 剪枝
		if v, ok := fm.miniTotalPrice[ct]; ok {
			for nk, p := range v {
				if nk >= k && p <= totalPrice {
					continue depthLoop
				}
			}
		}

		if _, ok := fm.miniTotalPrice[ct]; !ok {
			fm.miniTotalPrice[ct] = make(map[int]int)
		}
		fm.miniTotalPrice[ct][k] = totalPrice
		if ct == dst {
			continue
		}

		if k > 0 {
			fm.cacheSearchPath = append(fm.cacheSearchPath, visitedNodeInfo{ct, totalPrice})
			fm.calculateCheapestPrice(ct, dst, k-1)
			fm.cacheSearchPath = fm.cacheSearchPath[:len(fm.cacheSearchPath)-1]
		}
	}
}

func testCheapestFlightWithinKStops() {
	fmt.Println(cheapestFlightWithinKStopsDijkstra(0, [][]int{{0, 1, 100}, {1, 2, 100}, {1, 3, 600}, {2, 3, 200}, {2, 0, 100}}, 0, 3, 1) == 700)
}

func cheapestFlightWithinKStopsDijkstra(n int, flights [][]int, src, dst, k int) int {
	graph := make([]WeightEdge, 0)
	for _, route := range flights {
		graph = append(graph, WeightEdge{route[0], route[1], 1})
	}

	withinK := Dijkstra(graph, src)
	citistWithinK := make(map[int]struct{})
	for ct, dis := range withinK {
		if dis <= k+1 {
			citistWithinK[ct] = struct{}{}
		}
	}

	if _, ok := citistWithinK[dst]; !ok {
		return -1
	}

	newGraph := make([]WeightEdge, 0)
	for _, route := range flights {
		if _, ok := citistWithinK[route[0]]; ok {
			if _, ok := citistWithinK[route[1]]; ok {
				newGraph = append(newGraph, WeightEdge{route[0], route[1], route[2]})
			}
		}
	}
	rc := Dijkstra(newGraph, src)
	return rc[dst]
}

package day08

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j     int
	distance float64
}

type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.size[rootX] < uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

func (uf *UnionFind) GetCircuitSizes() []int {
	circuits := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		circuits[root] = uf.size[root]
	}

	sizes := make([]int, 0, len(circuits))
	for _, size := range circuits {
		sizes = append(sizes, size)
	}
	return sizes
}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day08")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	points := parsePoints(lines)

	if part == 1 {
		result := connectJunctionBoxes(points, 1000)
		fmt.Printf("Part 1:%d\n", result)
	} else {
		result := findLastConnection(points)
		fmt.Printf("Part 2: %d\n", result)
	}
}

func parsePoints(lines []string) []Point {
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z})
	}
	return points
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func connectJunctionBoxes(points []Point, connections int) int {
	n := len(points)

	edgeHeap := &EdgeHeap{}
	heap.Init(edgeHeap)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			heap.Push(edgeHeap, Edge{i, j, dist})
		}
	}

	uf := NewUnionFind(n)
	attempts := 0

	for edgeHeap.Len() > 0 && attempts < connections {
		edge := heap.Pop(edgeHeap).(Edge)
		uf.Union(edge.i, edge.j)
		attempts++
	}

	sizes := uf.GetCircuitSizes()
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}

	return result
}

func findLastConnection(points []Point) int {
	n := len(points)

	edgeHeap := &EdgeHeap{}
	heap.Init(edgeHeap)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			heap.Push(edgeHeap, Edge{i, j, dist})
		}
	}

	uf := NewUnionFind(n)
	numCircuits := n
	var lastI, lastJ int

	for edgeHeap.Len() > 0 && numCircuits > 1 {
		edge := heap.Pop(edgeHeap).(Edge)
		if uf.Union(edge.i, edge.j) {
			numCircuits--
			lastI = edge.i
			lastJ = edge.j
		}
	}

	return points[lastI].x * points[lastJ].x
}

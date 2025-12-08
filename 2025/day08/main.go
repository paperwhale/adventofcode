package main

import (
	"bufio"
	"container/heap"
	_ "embed"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/paperwhale/adventofcode/internal/must"
)

//go:embed input.txt
var input string

type Point struct {
	X int
	Y int
	Z int
}

type Pair struct {
	ID1      int
	ID2      int
	Distance float64
}

// MinHeap is a min-heap of Pair values.
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	pair := old[n-1]
	*h = old[0 : n-1]
	return pair
}

// UnionFind represents the disjoint set data structure
type UnionFind struct {
	parent []int
	size   []int
	count  int
}

// NewUnionFind creates a new UnionFind structure with n elements (0 to n-1).
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}
	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// Find determines the representative (root) of the set that element x belongs to.
// Uses path compression for efficiency.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing elements x and y.
// Uses union by rank to keep the tree shallow.
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return
	}
	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}
	uf.count--
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) ComponentSizes() []int {
	var sizes []int
	for i := 0; i < len(uf.parent); i++ {
		if uf.parent[i] == i {
			sizes = append(sizes, uf.size[i])
		}
	}
	return sizes
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(strings.NewReader(input), 1000))
	fmt.Println(Part2(strings.NewReader(input)))
}

func Part1(r io.Reader, numConnections int) int {
	points := parseInput(r)
	n := len(points)
	distances := make(MinHeap, 0, n*(n-1)/2)
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			distances = append(distances, Pair{
				ID1:      i,
				ID2:      j,
				Distance: distance(points[i], points[j]),
			})
		}
	}
	heap.Init(&distances)
	uf := NewUnionFind(len(points))
	for _ = range numConnections {
		pair := heap.Pop(&distances).(Pair)
		uf.Union(pair.ID1, pair.ID2)
	}
	sizes := uf.ComponentSizes()
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})
	return sizes[0] * sizes[1] * sizes[2]
}

func Part2(r io.Reader) int {
	points := parseInput(r)
	n := len(points)
	distances := make(MinHeap, 0, n*(n-1)/2)
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			distances = append(distances, Pair{
				ID1:      i,
				ID2:      j,
				Distance: distance(points[i], points[j]),
			})
		}
	}
	heap.Init(&distances)
	uf := NewUnionFind(len(points))
	result := -1
	for len(distances) > 0 {
		pair := heap.Pop(&distances).(Pair)
		uf.Union(pair.ID1, pair.ID2)
		if uf.Count() == 1 {
			result = points[pair.ID1].X * points[pair.ID2].X
			break
		}
	}
	return result
}

func parseInput(r io.Reader) []Point {
	var points []Point
	s := bufio.NewScanner(r)
	for s.Scan() {
		p := strings.Split(s.Text(), ",")
		must.Get(strconv.Atoi(p[0]))
		points = append(points, Point{
			X: must.Get(strconv.Atoi(p[0])),
			Y: must.Get(strconv.Atoi(p[1])),
			Z: must.Get(strconv.Atoi(p[2])),
		})
	}
	return points
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	dz := float64(p1.Z - p2.Z)

	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2))
}

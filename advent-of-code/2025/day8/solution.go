package main

import (
	"aoc/util"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"
)

type point struct {
	x int
	y int
	z int
}

type distance struct {
	p1       int // index of the point in the points array
	p2       int
	distance float64
}

type unionFind struct {
	parent []int
	size   []int
}

func newUnionFind(n int) *unionFind {
	p := make([]int, n)
	s := make([]int, n)
	for i := range n {
		p[i] = i
		s[i] = 1
	}
	return &unionFind{p, s}
}

func (uf *unionFind) find(a int) int {
	if uf.parent[a] == a {
		return a
	}
	return uf.find(uf.parent[a])
}

func (uf *unionFind) union(a, b int) int {
	i := uf.find(a)
	j := uf.find(b)
	if i == j {
		return 0
	}
	if uf.size[i] > uf.size[j] {
		uf.parent[j] = i
		uf.size[i] += uf.size[j]
	} else {
		uf.parent[i] = j
		uf.size[j] += uf.size[i]
	}
	return 1
}

func euclideanDistance(point1, point2 point) float64 {
	d1 := float64(point1.x - point2.x)
	d2 := float64(point1.y - point2.y)
	d3 := float64(point1.z - point2.z)
	return math.Sqrt(d1*d1 + d2*d2 + d3*d3)
}

func solvePartI(lines [][]byte, limit int) {
	n := len(lines)
	points := make([]point, n)
	for i, line := range lines {
		pstr := strings.Split(string(line), ",")
		x := util.GetInt(pstr[0])
		y := util.GetInt(pstr[1])
		z := util.GetInt(pstr[2])
		points[i] = point{x, y, z}
	}

	distances := make([]distance, 0, n*(n-1)/2)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := distance{i, j, euclideanDistance(points[i], points[j])}
			distances = append(distances, d)
		}
	}

	slices.SortFunc(distances, func(a, b distance) int {
		return cmp.Compare(a.distance, b.distance)
	})

	uf := newUnionFind(n)
	for _, d := range distances[:limit] {
		uf.union(d.p1, d.p2)
	}

	sizes := slices.Clone(uf.size)
	slices.Sort(sizes)
	ans := 1
	for _, s := range sizes[len(sizes)-3:] {
		ans *= s
	}
	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	n := len(lines)
	points := make([]point, n)
	for i, line := range lines {
		pstr := strings.Split(string(line), ",")
		x := util.GetInt(pstr[0])
		y := util.GetInt(pstr[1])
		z := util.GetInt(pstr[2])
		points[i] = point{x, y, z}
	}

	distances := make([]distance, 0, n*(n-1)/2)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := distance{i, j, euclideanDistance(points[i], points[j])}
			distances = append(distances, d)
		}
	}

	slices.SortFunc(distances, func(a, b distance) int {
		return cmp.Compare(a.distance, b.distance)
	})

	uf := newUnionFind(n)
	unions := 0
	ans := 0
	for _, d := range distances {
		unions += uf.union(d.p1, d.p2)
		if unions == n-1 {
			ans = points[d.p1].x * points[d.p2].x
			break
		}
	}
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 8 Test ----\n")
	solvePartI(testLines, 10) // 40
	solvePartII(testLines)    // 25272
	fmt.Printf("---- Day 8 Input ----\n")
	solvePartI(inputLines, 1000) // 115885
	solvePartII(inputLines)      // 274150525
}

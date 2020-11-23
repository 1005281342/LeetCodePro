package main

import (
	"sort"
)

type Int2Slice [][]int

func (p Int2Slice) Len() int           { return len(p) }
func (p Int2Slice) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p Int2Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int2Slice) Sort() { sort.Sort(p) }

func findMinArrowShots(points [][]int) int {

	if len(points) <= 1 {
		return len(points)
	}

	Int2Slice(points).Sort()
	//fmt.Println(points)

	var cnt int
	var tmp = points[0]
	for i := 1; i < len(points); i++ {
		// 不能合并
		if points[i][0] > tmp[1] {
			cnt++
			tmp = points[i]
			continue
		}
		// 可以合并
		tmp[0] = max(tmp[0], points[i][0])
		tmp[1] = min(tmp[1], points[i][1])
	}
	cnt++
	return cnt
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

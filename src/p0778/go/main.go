package main

import (
	"container/heap"
	"math"
)

var move = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// 堆
type NodeSlice []Node

type Node struct {
	x int
	y int
	d int
}

func NewNodeSlice() *NodeSlice {
	var ns NodeSlice
	return &ns
}

func (ns NodeSlice) Len() int {
	return len(ns)
}

func (ns NodeSlice) Less(i int, j int) bool {
	return ns[i].d < ns[j].d
}

func (ns NodeSlice) Swap(i int, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns *NodeSlice) Push(x interface{}) {
	*ns = append(*ns, x.(Node))
}

func (ns *NodeSlice) Pop() interface{} {
	old := *ns
	n := len(old)
	x := old[n-1]
	*ns = old[0 : n-1]
	return x
}

func swimInWater(heights [][]int) int {
	var (
		pq    = NewNodeSlice()
		node  Node
		a     int
		b     int
		dList [][]int
		midd  int
	)
	heap.Push(pq, Node{x: 0, y: 0, d: 0})

	// 最小体力记录
	dList = make([][]int, len(heights))
	for i := 0; i < len(heights); i++ {
		dList[i] = make([]int, len(heights[0]))
		for j := 0; j < len(heights[0]); j++ {
			dList[i][j] = math.MaxInt32
		}
	}
	dList[0][0] = heights[0][0]

	for pq.Len() > 0 {
		node = heap.Pop(pq).(Node)

		for i := 0; i < 4; i++ {
			a = move[i][0] + node.x
			b = move[i][1] + node.y
			// 越界
			if !(a >= 0 && a < len(heights) && b >= 0 && b < len(heights[0])) {
				continue
			}

			// 体力取最大的
			midd = max(node.d, heights[a][b], dList[node.x][node.y])
			// 如果这条路径的所需要的体力小于之前的
			if midd < dList[a][b] {
				dList[a][b] = midd
				heap.Push(pq, Node{x: a, y: b, d: dList[a][b]})
			}
		}
	}

	return max(dList[len(dList)-1][len(dList[0])-1], heights[len(dList)-1][len(dList[0])-1])
}

// max 最大值
func max(a int, nums ...int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			a = nums[i]
		}
	}

	return a
}

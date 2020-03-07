package main

var MOVE = [4][2]int{
	[2]int{1, 0},
	[2]int{-1, 0},
	[2]int{0, 1},
	[2]int{0, -1},
}

func orangesRotting(grid [][]int) int {
	var c int
	a := len(grid)
	b := len(grid[0])
	n := a * b
	dq := make([][3]int, 0, n)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if grid[i][j] == 2 {
				dq = append(dq, [3]int{i, j, 0})
			}
		}
	}
	for len(dq) > 0 {
		i, j := dq[0][0], dq[0][1]
		c = dq[0][2]
		for _, v := range MOVE {
			x, y := v[0], v[1]
			x += i
			y += j
			if 0 <= x && x < a && 0 <= y && y < b && grid[x][y] == 1 {
				dq = append(dq, [3]int{x, y, c + 1})
				grid[x][y] = 2
			}
		}
		dq = dq[1:]
	}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return c
}

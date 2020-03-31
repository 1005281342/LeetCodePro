package main

var MOVE = [4][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}

func surfaceArea(grid [][]int) int {

	length := len(grid)

	ans := 0

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if grid[i][j] <= 0 {
				continue
			}
			ans += 2
			for _, v := range MOVE {
				ni, nj := i+v[0], j+v[1]
				var nval int
				if 0 <= ni && ni < length && 0 <= nj && nj < length {
					nval = grid[ni][nj]
				}
				if grid[i][j] > nval {
					ans += grid[i][j] - nval
				}
			}
		}
	}
	return ans
}

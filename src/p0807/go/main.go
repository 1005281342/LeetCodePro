package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	length := len(grid)
	if length <= 1 {
		return 0
	}
	x := make([]int, length, length)
	for i, v := range grid {
		x[i] = arrayMaxNum(v...)
	}

	y := make([]int, length, length)
	for i := 0; i < length; i++ {
		t := make([]int, length, length)
		for j := 0; j < length; j++ {
			t[j] = grid[j][i]
		}
		y[i] = arrayMaxNum(t...)
	}

	ans := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			m := min(x[i], y[j])
			if m <= grid[i][j] {
				continue
			}
			ans += m - grid[i][j]
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func arrayMaxNum(a ...int) int {
	ans := 0
	for _, v := range a {
		if ans < v {
			ans = v
		}
	}
	return ans
}

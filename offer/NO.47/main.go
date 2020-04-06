package main

func maxValue(grid [][]int) int {

	x := len(grid)
	y := len(grid[0])
	var dp [][]int
	dp = make([][]int, x+1, x+1)
	for i :=0 ;i <=x; i++ {
		dp[i] = make([]int, y+1, y+1)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[x][y]

	// 方法1:
	//x := len(grid)
	//y := len(grid[0])
	//var dp = make([]int, y+1, y+1)
	//
	//for i := 1; i <= x; i++ {
	//	for j := 1; j <= y; j++ {
	//		dp[j] = max(dp[j], dp[j-1])+grid[i-1][j-1]
	//	}
	//}
	//return dp[y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

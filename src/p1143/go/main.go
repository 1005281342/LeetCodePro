package main

func longestCommonSubsequence(text1 string, text2 string) int {
	x := len(text1)
	y := len(text2)

	// 初始化
	var dp [][]int
	dp = make([][]int, x+1, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1, y+1)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

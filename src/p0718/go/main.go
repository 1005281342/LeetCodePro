package main

func findLength(A []int, B []int) int {

	x := len(A)
	y := len(B)
	ans := 0
	var dp [][]int
	dp = make([][]int, x+1, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1, y+1)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
				if dp[i+1][j+1] > ans {
					ans = dp[i+1][j+1]
				}
			} else {
				dp[i+1][j+1] = 0
			}
		}
	}
	return ans
}

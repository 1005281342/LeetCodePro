package main

const MOD = 1e9 + 7

func waysToStep(n int) int {
	dp := make([]int, n+1, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1]%MOD + dp[i-2]%MOD + dp[i-3]%MOD
	}
	return dp[n]
}

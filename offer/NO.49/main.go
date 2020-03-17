package main

func nthUglyNumber(n int) int {
	dp := make([]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	var a, b, c int
	for i := 1; i < n; i++ {
		dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)
		if dp[i] == dp[a]*2 {
			a += 1
		}
		if dp[i] == dp[b]*3 {
			b += 1
		}
		if dp[i] == dp[c]*5 {
			c += 1
		}
	}
	return dp[n-1]
}

func min(c ...int) int {
	ans := c[0]
	for _, v := range c[1:] {
		if v < ans {
			ans = v
		}
	}
	return ans
}

/*
class Solution:
    def nthUglyNumber(self, n: int) -> int:
        dp = [1 for _ in range(n)]
        idx_a = 0
        idx_b = 0
        idx_c = 0
        for i in range(1,n):
            dp[i] = min(dp[idx_a] * 2, dp[idx_b] * 3, dp[idx_c] * 5)
            if dp[i] == dp[idx_a] * 2:
                idx_a += 1
            if dp[i] == dp[idx_b] * 3:
                idx_b += 1
            if dp[i] == dp[idx_c] * 5:
                dp[i] == dp[idx_c] * 5
                idx_c += 1
        return dp[-1]
*/

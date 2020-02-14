package main

const MD = 1000000007

var ans [101]int

func init() {
	ans[0] = 1
	ans[1] = 1
	for i := 2; i <= 100; i++ {
		ans[i] = (ans[i-1] + ans[i-2]) % MD
	}
}

func numWays(n int) int {
	return ans[n]
}

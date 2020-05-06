package main

func mincostTickets(days []int, costs []int) int {

	dp := make([]int, days[len(days)-1]+1, days[len(days)-1]+1)

	index := 0
	i := 0
	for index < len(days) {
		i += 1
		if i != days[index] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = min(dp[max(i-1, 0)]+costs[0],
			dp[max(i-7, 0)]+costs[1],
			dp[max(i-30, 0)]+costs[2],
		)
		index += 1
	}
	return dp[len(days)-1]
}

func min(args ...int) int {
	ma := (1 << 31) - 1
	for _, v := range args {
		if v < ma {
			ma = v
		}
	}
	return ma
}

func max(args ...int) int {
	ma := -1
	for _, v := range args {
		if v > ma {
			ma = v
		}
	}
	return ma
}

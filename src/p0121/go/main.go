package main

func maxProfit(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > ans {
				ans = prices[j] - prices[i]
			}
		}
	}
	return ans
}

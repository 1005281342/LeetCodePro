package main


func maxProfit(prices []int) int {

	if len(prices) <= 1{
		return 0
	}
	// 股票最大利润
	m := 0
	// 可买入的最小股票值
	gm := prices[0]
	for i:=1 ;i < len(prices); i++ {
		m = max(m, prices[i]-gm)
		gm = min(gm, prices[i])
	}
	return m
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int  {
	if a < b {
		return a
	}
	return b
}
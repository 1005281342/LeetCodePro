package main

import "math"

func printNumbers(n int) []int {
	var cnt int
	cnt = int(math.Pow(10, float64(n)))
	ans := make([]int, 0, cnt)
	for i := 1; i < cnt; i++ {
		ans = append(ans, i)
	}
	return ans
}

package main

func climbStairs(n int) int {
	var (
		a, b = 1, 1
	)

	for i := 2; i < n; i++ {
		b, a = a+b, b
	}
	return b
}

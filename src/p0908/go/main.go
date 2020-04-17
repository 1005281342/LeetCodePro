package main

func smallestRangeI(A []int, K int) int {

	if len(A) <= 1 {
		return 0
	}

	mi := 10001
	ma := -1
	for _, v := range A {
		if v > ma {
			ma = v
		}
		if v < mi {
			mi = v
		}
	}

	d := ma - mi
	return max(d-2*K, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

func findSwapValues(array1 []int, array2 []int) []int {

	a := sum(array1...)
	b := sum(array2...)
	if a > b {
		a, b = b, a
	}
	c := b - a
	if c&1 == 1 {
		return []int{}
	}
	c /= 2

	ma := make(map[int]bool)
	for _, v := range array1 {
		ma[v] = true
	}

	mb := make(map[int]bool)
	for _, v := range array2 {
		mb[v] = true
	}

	for v1 := range ma {
		for v2 := range mb {
			if v1 > v2 && v1-v2 == c {
				return []int{v1, v2}
			} else if v2-v1 == c {
				return []int{v1, v2}
			}
		}
	}
	return []int{}
}

func sum(a ...int) int {
	ans := 0
	for _, v := range a {
		ans += v
	}
	return ans
}

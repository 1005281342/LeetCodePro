package main

func exchange(nums []int) []int {
	length := len(nums)
	x := (length + 1) >> 1
	a := make([]int, 0, x)
	b := make([]int, 0, x)
	for _, v := range nums {
		if v&1 == 0 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	return append(b, a...)
}

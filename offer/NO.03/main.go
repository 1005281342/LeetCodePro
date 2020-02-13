package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]int, len(nums)-1)

	for _, v := range nums {
		if m[v] >= 1 {
			return v
		} else {
			m[v] = 1
		}
	}

	return 0
}

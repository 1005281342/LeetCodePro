package main

func majorityElement(nums []int) int {
	c := len(nums) >> 1
	tm := make(map[int]int, c)

	for _, v := range nums {
		tm[v] += 1
	}

	for _, v := range tm {
		if v > c {
			return v
		}
	}
	return 0
}

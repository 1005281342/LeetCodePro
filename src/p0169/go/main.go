package main

func majorityElement(nums []int) int {
	t := len(nums) / 2
	mp := make(map[int]int, t)
	for _, v := range nums {
		mp[v] += 1
		if mp[v] > t {
			return v
		}
	}
	return 0
}

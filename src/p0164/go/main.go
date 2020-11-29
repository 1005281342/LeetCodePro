package main

import "sort"

func maximumGap(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	var d = 0
	for i := 0; i < len(nums)-1; i++ {
		var nd = nums[i+1] - nums[i]
		if nd <= d {
			continue
		}
		d = nd
	}
	return d
}

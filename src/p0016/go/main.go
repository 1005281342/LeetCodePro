package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		left  int
		right int
		ans   = 10000
		sum   int
	)
	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			k := abs(target - sum)
			if abs(ans-target) > k {
				ans = sum
			}
			if target > sum {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

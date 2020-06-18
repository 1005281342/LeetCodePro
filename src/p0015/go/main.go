package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var mp = make(map[[3]int]struct{})
	var (
		left  int
		right int
		sum   int
	)

	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				mp[[3]int{nums[i], nums[left], nums[right]}] = struct{}{}
				left += 1
				right -= 1
			} else if sum > 0 {
				right -= 1
			} else {
				left += 1
			}
		}
	}

	var ans = make([][]int, 0)
	for k := range mp {
		ans = append(ans, append(make([]int, 0, 3), k[:]...))
	}
	return ans
}

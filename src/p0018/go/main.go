package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var (
		left  int
		right int
		sum   int
		mp    = make(map[[4]int]struct{})
		ans   = make([][]int, 0)
	)

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			left = j + 1
			right = len(nums) - 1
			for left < right {
				sum = nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					mp[[4]int{nums[i], nums[j], nums[left], nums[right]}] = struct{}{}
					left += 1
					right -= 1
				} else if sum < target {
					left += 1
				} else {
					right -= 1
				}
			}
		}
	}

	for k := range mp {
		ans = append(ans, append(make([]int, 0, 4), k[:]...))
	}
	return ans
}

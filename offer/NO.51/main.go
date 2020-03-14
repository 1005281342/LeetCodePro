package main

func reversePairs(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				ans++
			}
		}
	}
	return ans
}

package main

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if target < nums[right]+nums[left] {
			right -= 1
		} else if target > nums[right]+nums[left] {
			left += 1
		} else {
			return []int{left, right}
		}
	}
	return []int{0, 0}
}

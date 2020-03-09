package main

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

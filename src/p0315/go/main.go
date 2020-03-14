package main

func countSmaller(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums)-1; i++ {
		cnt := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				cnt += 1
			}
		}
		ans[i] = cnt
	}
	return ans
}

package main

func searchInsert(nums []int, target int) int {
	var index = binSearchLeft(nums, target)
	if nums[index] < target {
		index += 1
	}
	return index
}

func binSearchLeft(nums []int, target int) int {
	var (
		left  int
		right = len(nums) - 1
		mid   int
	)
	for left < right {
		mid = left + ((right - left) >> 1)
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

package main

func pivotIndex(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	//  前缀和
	var sums = make([]int, len(nums))
	sums[0] = nums[0]
	for i := 0; i < len(nums)-1; i++ {
		sums[i+1] = sums[i] + nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		if sums[i]-nums[i] == sums[len(nums)-1]-sums[i] {
			return i
		}
	}

	return -1
}

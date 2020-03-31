package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
	//fmt.Println(nums)
	return len(nums)
}

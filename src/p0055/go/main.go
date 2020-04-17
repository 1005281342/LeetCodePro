package main

/*
IN:
[2,3,1,1,4]
[3,2,1,0,4]
[1,2]
[1,2,0,1]

OUT:
true
false
true
true
*/
func canJump(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}

	if length == 2 {
		return nums[0] >= 1
	}

	c := 1
	for i := 2; i <= length; i++ {
		if nums[length-i] < c {
			c += 1
		} else {
			c = 1
		}
	}
	return nums[0] >= c
}

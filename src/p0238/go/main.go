package main

func productExceptSelf(nums []int) []int {
	var (
		a = make([]int, len(nums), len(nums))
		b = make([]int, len(nums), len(nums))
		c = make([]int, len(nums), len(nums))
	)

	a[0] = nums[0]
	b[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		a[i] = nums[i] * a[i-1]
		b[len(nums)-i-1] = nums[len(nums)-i-1] * b[len(nums)-i]
	}

	c[0] = b[1]
	c[len(nums)-1] = a[len(nums)-1-1]
	for i := 1; i < len(nums)-1; i++ {
		c[i] = a[i-1] * b[i+1]
	}
	return c
}

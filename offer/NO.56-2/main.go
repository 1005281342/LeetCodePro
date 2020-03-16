package main

func singleNumber(nums []int) int {
	var a, b int

	for _, v := range nums {
		a = (a ^ v) & (^b)
		b = (b ^ v) & (^a)
	}
	return a
}

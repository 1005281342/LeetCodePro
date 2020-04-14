package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	nums := []int{2, 5, 16, 2, 6, 7, 2, 1, 2, 4, 5, 3}
	sorted(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// 就地快排
func sorted(nums []int, start, end int) {
	if start >= end {
		return
	}
	t := nums[end]
	left := 0
	right := end - 1
	for left <= right {

		// 当前值小于基准值且在基准值左边，左游标右移
		for left <= right && nums[left] < t {
			left += 1
		}
		// 当前值大于基准值切在基准值右边，右游标左移
		for left <= right && nums[right] > t {
			right -= 1
		}
		// 存在基准值左边的数大于基准值右边的数时，交换这两个数
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left += 1
			right -= 1
		}
	}

	// 将基准值交换到low、high之间，此时low, mid, high分离完成
	/*
	low: 小于基准值部分
	mid: 等于基准值部分
	high: 大于基准值部分
	 */
	nums[left], nums[end] = nums[end], nums[left]
	sorted(nums, start, left-1)
	sorted(nums, left+1, end)
}

//func sorted(nums []int) []int {
//	length := len(nums)
//	if length <= 1 {
//		return nums
//	}
//
//	// 从数组中随机取得一个基准值
//	t := nums[rand.Intn(length)]
//
//	// 小于基准值
//	low := make([]int, 0)
//	// 等于基准值
//	mid := make([]int, 0)
//	// 大于基准值
//	high := make([]int, 0)
//
//	for i := 0; i < length; i++ {
//		if nums[i] < t {
//			low = append(low, nums[i])
//		} else if nums[i] > t {
//			high = append(high, nums[i])
//		} else {
//			mid = append(mid, nums[i])
//		}
//	}
//
//	low = sorted(low)
//	high = sorted(high)
//	return append(append(low, mid...), high...)
//}

package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 5, 3, 3, 6, 9, 9, 6, 3, 7, 8, 2}
	fmt.Println(sorted(nums, 9))
}

// 适用在数组值在一定范围的数组排序
// 利用一个计数数组进行统计
func sorted(nums []int, maxN int) []int {
	dd := make([]int, maxN+1, maxN+1)
	for _, v := range nums {
		dd[v] += 1
	}
	ans := make([]int, 0, len(nums))
	for i, v := range dd {
		for v > 0 {
			ans = append(ans, i)
			v -= 1
		}
	}
	return ans
}
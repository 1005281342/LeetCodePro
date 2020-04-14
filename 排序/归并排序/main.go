package main

import "fmt"

func main() {
	var a = []int{1,4,5,2,3,8,5,2,8,33,15}
	fmt.Println(sorted(a))
}

func sorted(nums []int) []int {
	// 切片长度
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 当切分到可比较时
	if length == 2 && nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}

	// 切分与归并
	return merge(sorted(nums[:length>>1]), sorted(nums[length>>1:]))
}

// 归并
func merge(a, b []int) []int {
	var i, j int
	ans := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i += 1
		} else {
			ans = append(ans, b[j])
			j += 1
		}
	}
	if i < len(a) {
		ans = append(ans, a[i:]...)
	}
	if j < len(b) {
		ans = append(ans, b[j:]...)
	}
	return ans
}
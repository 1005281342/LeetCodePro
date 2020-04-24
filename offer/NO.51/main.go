package main

import "fmt"

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{4, 5, 6, 7}))
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
	fmt.Println(reversePairs([]int{1, 1, 1, 1, 1}))
}

func reversePairs(nums []int) int {

	//ans := 0
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := i; j < len(nums); j++ {
	//		if nums[i] > nums[j] {
	//			ans++
	//		}
	//	}
	//}
	//return ans
	mergeSort(nums)
	t := cnt
	cnt = 0
	return t
}

var cnt int

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 当切分到可直接比较时
	if length == 2 && nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
		cnt += 1
	}
	mid := length >> 1
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(a, b []int) []int {
	ans := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			//fmt.Println(cnt, a[i], b[j])
			cnt += len(a) - i
			ans = append(ans, b[j])
			j += 1
		} else if a[i] < b[j] {
			ans = append(ans, a[i])
			i += 1
		} else {
			ans = append(ans, a[i])
			i += 1
			ans = append(ans, b[j])
			ii := i
			for ; ii < len(a); ii++ {
				if a[ii] > b[j] {
					break
				}
			}
			cnt += max(len(a)-ii, 0)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

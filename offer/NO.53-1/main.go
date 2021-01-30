package main

func search(nums []int, target int) int {
	var cnt int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			cnt++
		}
	}

	return cnt
}

//// 找到目标值所在的位置
//func binSearch(nums []int, left, right, target int) int {
//
//	if left == right {
//		if nums[left] == target {
//			return left
//		} else {
//			return -1 // 目标值不存在
//		}
//	}
//
//	if left > right {
//		return -1
//	}
//
//	mid := left + (right-left)>>1 // 中间索引
//	if nums[mid] < target {       // 目标值在右边
//		return binSearch(nums, mid+1, right, target)
//	} else if nums[mid] > target {
//		return binSearch(nums, left, mid-1, target)
//	} else {
//		return mid
//	}
//}
//
//func search(nums []int, target int) int {
//	/*
//		排序的数组可以使用二分查找找到它所在位置
//		然后向左右去查找并统计
//	*/
//	index := binSearch(nums, 0, len(nums)-1, target)
//	if index < 0 {
//		return 0
//	}
//
//	cnt := 1
//	// 往左查找
//	a := index - 1
//	for a >= 0 {
//		if nums[a] != target {
//			break
//		}
//		cnt += 1
//		a -= 1
//	}
//	// 往右查找
//	a = index + 1
//	for a < len(nums) {
//		if nums[a] != target {
//			break
//		}
//		cnt += 1
//		a += 1
//	}
//	return cnt
//}

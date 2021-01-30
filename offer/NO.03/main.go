package main

func findRepeatNumber(nums []int) int {
	var mp = make(map[int]struct{})
	for i := 0; i < len(nums); i++ {

		if _, exist := mp[nums[i]]; exist {
			return nums[i]
		}

		mp[nums[i]] = struct{}{}
	}

	return 0
}

//func findRepeatNumber(nums []int) int {
//	m := make(map[int]int, len(nums)-1)
//
//	for _, v := range nums {
//		if m[v] >= 1 {
//			return v
//		} else {
//			m[v] = 1
//		}
//	}
//
//	return 0
//}

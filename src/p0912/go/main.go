package main

import "math/rand"

func sortArray(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	// low, mid, high
	// 每次随机一个数组下标取得 参照值
	tmp := arr[rand.Intn(length)]
	//tmp := arr[time.Now().Unix() % int64(length)]
	low := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	high := make([]int, 0, 0)
	//mid = append(mid, tmp)

	for i := 0; i < length; i++ {
		if arr[i] < tmp {
			low = append(low, arr[i])
		} else if arr[i] > tmp {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = sortArray(low), sortArray(high)
	myArr := append(append(low, mid...), high...)
	return myArr
}

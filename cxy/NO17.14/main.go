package main

import (
	"math/rand"
)

func main() {
	smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4)
}

func smallestK(arr []int, k int) []int {
	arr = quickSort(arr)
	//fmt.Println(arr)
	return arr[:k]
}

func quickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	index := rand.Intn(length)
	t := arr[index]
	low := make([]int, 0)
	mid := make([]int, 0)
	high := make([]int, 0)

	for i := 0; i < length; i++ {
		if arr[i] < t {
			low = append(low, arr[i])
		} else if arr[i] > t {
			high = append(high, arr[i])
		} else {
			mid = append(mid, t)
		}
	}

	return append(append(quickSort(low), mid...), quickSort(high)...)
}

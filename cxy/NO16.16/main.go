package main

func subSort(array []int) []int {

	max, min := -1000000, 1000000
	left, right := -1, -1
	for i := 0; i < len(array); i++ {
		if array[i] < max {
			right = i
		} else {
			max = array[i]
		}
	}
	for j := len(array) - 1; j >= 0; j-- {
		if array[j] > min {
			left = j
		} else {
			min = array[j]
		}
	}

	if left == -1 || right == -1 {
		return []int{}
	}

	return []int{left, right}
}

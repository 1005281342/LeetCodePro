package main

func minArray(numbers []int) int {

	length := len(numbers)
	if length == 1 {
		return numbers[0]
	}
	ans := numbers[0]
	for i := 1; i < length; i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return ans
}

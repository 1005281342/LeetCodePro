package main

func findContinuousSequence(target int) [][]int {

	ans := make([][]int, 0, target>>2)
	i, j, cnt := 1, 2, 3

	for i <= target/2 {
		if cnt < target {
			j += 1
			cnt += j
		} else if cnt > target {
			cnt -= i
			i += 1
		} else {
			t := make([]int, 0, j-i+1)
			for a := i; a <= j; a++ {
				t = append(t, a)
			}
			ans = append(ans, t)
			j += 1
			cnt += j
		}
	}
	return ans
}

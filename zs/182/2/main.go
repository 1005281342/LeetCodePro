package main

func numTeams(rating []int) int {
	length := len(rating)
	if length < 3 {
		return 0
	}
	ans := 0
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if (rating[i] > rating[j] && rating[j] > rating[k]) || (rating[i] < rating[j] && rating[j] < rating[k]) {
					ans += 1
				}
			}
		}
	}
	return ans
}

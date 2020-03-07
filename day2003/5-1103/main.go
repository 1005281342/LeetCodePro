package main

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people, num_people)
	cnt := 0
	for candies > (cnt + 1) {
		ans[cnt%num_people] += cnt + 1
		candies -= cnt + 1
		cnt += 1
	}
	ans[cnt%num_people] += candies
	return ans
}

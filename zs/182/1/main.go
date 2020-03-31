package main

func findLucky(arr []int) int {
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v] += 1
	}
	ans := -1
	for k, v := range mp {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}

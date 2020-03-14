package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	ans := 0
	for i, v := range informTime {
		if v == 0 {
			k := helper(i, manager, informTime)
			if k > ans {
				ans = k
			}
		}
	}
	return ans
}

func helper(i int, manager []int, informTime []int) int {
	ans := 0
	k := i
	for {
		if manager[k] == -1 {
			return ans
		}
		ans += informTime[manager[k]]
		k = manager[k]
	}
}

package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	length1 := len(word1)
	length2 := len(word2)

	if length2 == 0 || length1 == 0 {
		return max(length1, length2)
	}

	var dp [][]int
	dp = make([][]int, length1+1, length1+1)
	for i := 0; i <= length1; i++ {
		dp[i] = make([]int, length2+1, length2+1)
	}

	for i := 0; i <= length1; i++ {
		for j := 0; j <= length2; j++ {
			dp[i][j] = i + j
		}
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			var t int
			if word1[i-1] != word2[j-1] {
				t = 1
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+t)
		}
	}
	fmt.Println(dp)
	return dp[length1][length2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a ...int) int {
	var ans = 1<<31 - 1
	for _, c := range a {
		if c < ans {
			ans = c
		}
	}
	return ans
}

package main

const (
	a   = 1 << iota
	e   = 1 << iota
	i   = 1 << iota
	o   = 1 << iota
	u   = 1 << iota
	inf = 1 << 31
)

var mp = map[byte]int{
	'a': a,
	'e': e,
	'i': i,
	'o': o,
	'u': u,
}

func findTheLongestSubstring(s string) int {
	var (
		dp    []int
		index int
		ans   int
	)
	dp = make([]int, 32, 32)
	dp[0] = -1
	for i := 1; i < 32; i++ {
		dp[i] = -inf
	}

	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			index ^= mp[s[i]]
		}

		if dp[index] != -inf {
			ans = max(ans, i-dp[index])
		} else {
			dp[index] = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

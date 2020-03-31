package main

func longestPalindrome(s string) int {
	mp := make(map[int32]int, len(s))
	for _, v := range []rune(s) {
		mp[v] += 1
	}

	ans := 0
	j := 0
	for _, cnt := range mp {
		ans += cnt
		if cnt&1 == 1 {
			ans -= 1
			j = 1
		}
	}
	return ans + j
}
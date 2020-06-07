package main

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	if len(s) == 2 && s[0] == s[1] {
		return s
	}

	var ans string
	for a := 0; a < len(s)-1; a++ {
		// 中心扩散法
		// 奇数
		i := a - 1
		j := a + 1
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i -= 1
			j += 1
		}
		if j-i-1 > len(ans) {
			ans = s[i+1 : j]
		}

		i = a
		j = a + 1
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i -= 1
			j += 1
		}
		if j-i-1 > len(ans) {
			ans = s[i+1 : j]
		}
	}

	return ans
}

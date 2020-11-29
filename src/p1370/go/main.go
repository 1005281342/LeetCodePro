package main

var charList = make([]int, 26)
var ans []byte

func sortString(s string) string {
	for i := range s {
		charList[s[i]-'a']++
	}

	for len(ans) < len(s) {
		for c := byte('a'); c <= 'z'; c++ {
			if charList[c-'a'] > 0 {
				ans = append(ans, c)
				charList[c-'a']--
			}
		}
		for c := byte('z'); c >= 'a'; c-- {
			if charList[c-'a'] > 0 {
				ans = append(ans, c)
				charList[c-'a']--
			}
		}
	}
	charList = make([]int, 26)
	var t = string(ans)
	ans = []byte{}
	return t
}

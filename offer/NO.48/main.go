package main

func lengthOfLongestSubstring(s string) int {
	ans := 0
	mp := make(map[uint8]int)
	i := 0
	for i < len(s) {
		k := s[i]
		if _, ok := mp[k]; ok {
			if len(mp) > ans {
				ans = len(mp)
			}
			i = mp[k] + 1
			mp = make(map[uint8]int)
			continue
		}
		mp[k] = i
		i += 1
	}
	if len(mp) > ans {
		ans = len(mp)
	}
	return ans
}

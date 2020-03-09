package main

func firstUniqChar(s string) byte {
	ans := byte(' ')
	mp := make(map[int32]int)
	k := []rune(s)
	for _, v := range k {
		mp[v] += 1
	}
	for _, v := range k {
		if mp[v] == 1 {
			return byte(v)
		}
	}
	return ans
}

package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1l := []byte(s1)
	s2l := []byte(s2)
	mp := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		mp[s1l[i]] += 1
		mp[s2l[i]] -= 1
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}

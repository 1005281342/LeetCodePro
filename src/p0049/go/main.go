package main

func groupAnagrams(strs []string) [][]string {
	var mp = make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		var kt = [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			kt[strs[i][j]-'a']++
		}
		if _, ok := mp[kt]; ok {
			mp[kt] = append(mp[kt], strs[i])
		} else {
			mp[kt] = []string{strs[i]}
		}
	}

	var ans = make([][]string, 0)
	for k := range mp {
		ans = append(ans, mp[k])
	}
	return ans
}

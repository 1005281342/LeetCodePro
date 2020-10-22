package main

func partitionLabels(S string) []int {

	if len(S) <= 0 {
		return []int{}
	}

	// 统计字母分布
	var mp = make(map[byte]int)
	for i := 0; i < len(S); i++ {
		mp[S[i]]++
	}

	var (
		index int

		cnt int
		tmp = make(map[byte]struct{})
		ans []int
	)

	for index < len(S) {

		if len(tmp) == 0 && cnt > 0 {
			ans = append(ans, cnt)
			cnt = 0
			tmp = make(map[byte]struct{})
		}

		var v = S[index]
		tmp[v] = struct{}{}
		mp[v]--
		cnt++
		if mp[v] == 0 {
			delete(mp, v)
			delete(tmp, v)
		}

		index++
	}
	if cnt > 0 {
		ans = append(ans, cnt)
	}
	return ans
}

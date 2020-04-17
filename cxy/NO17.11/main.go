package main

func findClosest(words []string, word1 string, word2 string) int {

	var a, b = -1, -1
	var ans = len(words)
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			a = i
			if b != -1 && abs(a-b) < ans {
				ans = abs(a - b)
			}
		}
		if words[i] == word2 {
			b = i
			if a != -1 && abs(b-a) < ans {
				ans = abs(b - a)
			}
		}
	}
	return ans

	//var mp = make(map[string][]int)
	//for i, v := range words {
	//	if _, ok := mp[v]; !ok {
	//		mp[v] = []int{i}
	//	} else {
	//		mp[v] = append(mp[v], i)
	//	}
	//}
	//
	//ans := len(words)
	//for _, a := range mp[word1] {
	//	for _, b := range mp[word2] {
	//		if abs(a-b) < ans {
	//			ans = abs(a-b)
	//		}
	//	}
	//}
	//return ans
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

package main

func countCharacters(words []string, chars string) int {
	k := make(map[int]map[int32]int, len(words)/3)
	for i, v := range words {
		t := make(map[int32]int, len(v)/3)
		for _, c := range v {
			t[c-97] += 1
		}
		k[i] = t
	}
	cs := make(map[int32]int, len(chars)/3)
	for _, c := range chars {
		cs[c-97] += 1
	}

	ans := 0
	for i := 0; i < len(words); i++ {
		flag := true
		for m, n := range k[i] {
			if n > cs[m] {
				flag = false
			}
		}
		if flag {
			ans += len(words[i])
		}
	}
	return ans
}

// 思路错误
//func countCharacters(words []string, chars string) int {
//	k := make(map[int]map[int32]int, len(words)/3)
//	for i, v := range words {
//		t := make(map[int32]int, len(v)/3)
//		for _, c := range v {
//			t[c-97] += 1
//		}
//		k[i] = t
//	}
//	cs := make(map[int32]int, len(chars)/3)
//	for _, c := range chars {
//		cs[c-97] += 1
//	}
//
//	minLen := len(words) * len(chars)
//	for i := 0; i < len(words)-1; i++ {
//		for j := i + 1; j < len(words); j++ {
//			flag := true
//			for m, n := range cs {
//				if k[i][m]+k[j][m] < n {
//					flag = !flag
//					break
//				}
//			}
//			if flag && len(words[i])+len(words[j]) < minLen {
//				minLen = len(words[i]) + len(words[j])
//			}
//		}
//	}
//	return minLen
//}

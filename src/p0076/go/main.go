package main

/*
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
*/

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}

func minWindow(s string, t string) string {

	if len(s) < len(t) || len(s) == 0 {
		return ""
	}

	var (
		tMap   = make(map[byte]int)
		cntMap = make(map[byte]int)

		left  int
		right int

		lAns   int = -1
		rAns   int = -1
		ansLen int = (1 << 31) - 1

		distance int
	)

	for i := 0; i < len(t); i++ {
		tMap[t[i]] += 1
	}
	//fmt.Println(tMap)

	for right < len(s) {

		cntMap[s[right]] += 1
		if v, ok := tMap[s[right]]; ok && v >= cntMap[s[right]] {
			distance += 1
		}

		for distance == len(t) && left <= right {
			var newLen = right - left + 1
			if newLen < ansLen {
				lAns = left
				rAns = right
				ansLen = newLen
			}

			cntMap[s[left]] -= 1
			if cntMap[s[left]] < tMap[s[left]] {
				distance -= 1
			}
			left += 1
		}

		right += 1
	}

	if lAns == -1 {
		return ""
	}

	return s[lAns : rAns+1]
}

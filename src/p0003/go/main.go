package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func lengthOfLongestSubstring(s string) int {

	// if len(s) <= 1{
	//     return len(s)
	// }

	var (
		left  int = 0
		right int = 0
		mp        = make(map[byte]int)
		ans   int
	)

	for i := 0; i < len(s); i++ {
		if v, ok := mp[s[i]]; !ok || v < left {
		} else {
			if right-left > ans {
				ans = right - left
			}
			left = v + 1
		}
		right += 1
		mp[s[i]] = i
		//fmt.Println(mp, left, right)
	}

	if right-left > ans {
		ans = right - left
	}

	return ans
}

package main

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	var mp = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var ans = make([]string, 0)
	ans = append(ans, mp[digits[0]]...)

	for i := 1; i < len(digits); i++ {
		length := len(ans)
		for _, v := range mp[digits[i]] {
			for j := 0; j < length; j++ {
				ans = append(ans, ans[j]+v)
			}
		}
		ans = ans[length:]
	}

	return ans

}

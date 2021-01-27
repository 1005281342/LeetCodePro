package main

import "strings"

func main() {

}

func reformatNumber(number string) string {
	var (
		nums []string
		mp   = map[string]struct{}{
			"0": {},
			"1": {},
			"2": {},
			"3": {},
			"4": {},
			"5": {},
			"6": {},
			"7": {},
			"8": {},
			"9": {},
		}
	)

	for i := 0; i < len(number); i++ {
		if _, ok := mp[number[i:i+1]]; ok {
			nums = append(nums, number[i:i+1])
		}
	}

	var (
		nl  = len(nums)
		ans string
	)
	if nl%3 == 0 {
		for i := 0; i < nl-3; i += 3 {
			ans += strings.Join(nums[i:i+3], "")
			ans += "-"
		}
		ans += strings.Join(nums[nl-3:nl], "")
		return ans
	}

	if (nl-2)%3 == 0 {
		for i := 0; i < nl-2; i += 3 {
			ans += strings.Join(nums[i:i+3], "")
			ans += "-"
		}
		ans += strings.Join(nums[nl-2:nl], "")
		return ans
	}

	for i := 0; i < nl-4; i += 3 {
		ans += strings.Join(nums[i:i+3], "")
		ans += "-"
	}
	ans += strings.Join(nums[nl-4:nl-2], "")
	ans += "-"
	ans += strings.Join(nums[nl-2:nl], "")
	return ans
}

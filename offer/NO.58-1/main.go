package main

import "strings"

func reverseWords(s string) string {
	ts := strings.Split(s, " ")
	tt := make([]string, 0, len(ts))
	for _, c := range ts {
		if c != "" {
			tt = append(tt, c)
		}
	}
	left := 0
	right := len(tt) -1
	for left < right {
		tt[left], tt[right] = tt[right], tt[left]
		left += 1
		right -= 1
	}
	return strings.Join(tt, " ")
}

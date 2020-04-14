package main

import "strings"

func reverseWords(s string) string {
	p := strings.Split(s, " ")
	var t = make([]string, 0)
	for _, v := range p {
		if len(v) > 0 {
			t = append(t, v)
		}
	}
	left := 0
	right := len(t) - 1
	for left < right {
		aa := t[right]
		bb := t[left]
		t[left], t[right] = strings.Replace(aa, " ", "", -1), strings.Replace(bb, " ", "", -1)
		left += 1
		right -= 1
	}
	return strings.Join(t, " ")
}

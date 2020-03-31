package main

import "fmt"

func isFlipedString(s1 string, s2 string) bool {

	if s1 == s2 {
		return true
	}
	for i := 0; i < len(s1); i++ {
		t := s1[i:] + s1[:i]
		if t == s2 {
			return true
		}
	}
	return false
}

func main() {
	ans := isFlipedString("1", "1")
	fmt.Println(ans)
}

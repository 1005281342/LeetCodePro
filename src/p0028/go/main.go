package main

func strStr(haystack string, needle string) int {

	if haystack == needle {
		return 0
	}

	var (
		offset = len(needle)
	)

	for i := 0; i < len(haystack)-offset; i++ {
		if haystack[i:i+offset] == needle {
			return i
		}
	}

	return -1
}

package main

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		len1, len2 = len2, len1
		str1, str2 = str2, str1
	}

	for i := len2; i > 0; i-- {
		p := len2 / i
		if float64(len2)/float64(i)-float64(p) > 0.000001 {
			continue
		}
		q := len1 / i
		if float64(len1)/float64(i)-float64(q) > 0.000001 {
			continue
		}
		if helper(str2, p, i) && helper(str1, q, i) && str1[:i] == str2[:i] {
			return str1[:i]
		}

	}
	return ""
}

func helper(s string, p, d int) bool {
	flag := true
	tmp := s[:d]
	for i := 1; i < p; i++ {
		if s[i*d:i*d+d] != tmp {
			flag = false
			break
		}
	}
	return flag
}

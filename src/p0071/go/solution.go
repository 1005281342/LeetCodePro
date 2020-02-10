package main

import "strings"

func simplifyPath(path string) string {
	plst := strings.Split(path, "/")

	ans := make([]string, 0, len(plst))
	for _, v := range plst {
		if v == ".." {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else if v == "." || v == "" {
			continue
		} else {
			ans = append(ans, v)
		}
	}

	return "/" + strings.Join(ans, "/")
}

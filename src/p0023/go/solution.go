package main

func isValid(s string) bool {
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	help := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if len(help) > 0 && m[string(s[i])] == help[len(help)-1] {
			help = help[:len(help)-1]
		} else {
			help = append(help, string(s[i]))
		}
	}
	return len(help) == 0
}

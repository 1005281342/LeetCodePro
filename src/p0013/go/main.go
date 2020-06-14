package main

import "fmt"

func main() {
	x := romanToInt("MCMXCIV")
	fmt.Println(x)
}

func romanToInt(s string) int {
	var mp = map[int]string{
		3000: "MMM", 2000: "MM", 1000: "M",
		900: "CM", 800: "DCCC", 700: "DCC", 600: "DC",
		500: "D", 400: "CD", 300: "CCC", 200: "CC", 100: "C",
		90: "XC", 80: "LXXX", 70: "LXX", 60: "LX",
		50: "L", 40: "XL", 30: "XXX", 20: "XX", 10: "X",
		9: "IX", 8: "VIII", 7: "VII", 6: "VI",
		5: "V", 4: "IV", 3: "III", 2: "II", 1: "I",
	}

	var mt = make(map[string]int)
	for k, v := range mp {
		mt[v] = k
	}
	var ans int
	var index int
	var start int
	for index < len(s) {
		if _, ok := mt[s[start:index+1]]; !ok {
			fmt.Println(ans, mt[s[start:index]], s[start:index])
			ans += mt[s[start:index]]
			start = index
		}
		index += 1
	}
	ans += mt[s[start:index]]
	return ans
}

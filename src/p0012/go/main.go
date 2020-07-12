package main

func intToRoman(num int) string {
	var mp = map[int]string{
		3000: "MMM", 2000: "MM", 1000: "M",
		900: "CM", 800: "DCCC", 700: "DCC", 600: "DC",
		500: "D", 400: "CD", 300: "CCC", 200: "CC", 100: "C",
		90: "XC", 80: "LXXX", 70: "LXX", 60: "LX",
		50: "L", 40: "XL", 30: "XXX", 20: "XX", 10: "X",
		9: "IX", 8: "VIII", 7: "VII", 6: "VI",
		5: "V", 4: "IV", 3: "III", 2: "II", 1: "I",
	}
	var cnt, t int
	var ans string
	for num > 0 {
		num, t = num/10, num%10
		ans = mp[t*pow(cnt, 10)] + ans
		cnt += 1
	}
	return ans
}

func pow(c int, v int) int {

	var cnt, x int
	x = 1
	for cnt < c {
		x *= v
		cnt += 1
	}
	return x
}

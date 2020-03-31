package main

func hasGroupsSizeX(deck []int) bool {
	mp := make(map[int]int)

	for _, v := range deck {
		mp[v] += 1
	}

	d := len(deck) + 1

	for _, v := range mp {
		if d == len(deck)+1 {
			d = v
		}
		d = gcd(d, v)
		if d < 2 {
			return false
		}
	}
	return d >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

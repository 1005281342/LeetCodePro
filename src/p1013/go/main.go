package main

func canThreePartsEqualSum(A []int) bool {
	cnt := 0
	for _, v := range A {
		cnt += v
	}
	if cnt%3 != 0 {
		return false
	}
	tk := cnt / 3
	tc := 0
	c := 0
	for _, v := range A{
		tc += v
		if tc == tk {
			tc = 0
			c += 1
		}
	}
	return c == 3 || (c >3 && tk == 0)
}

func main() {
	a := []int{10,-10,10,-10,10,-10,10,-10}
	canThreePartsEqualSum(a)
}
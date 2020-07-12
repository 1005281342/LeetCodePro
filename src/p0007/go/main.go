package main

func reverse(x int) int {

	var (
		flag   bool = x < 0
		t      int64
		list   = make([]int, 0)
		a      int
		length int
	)

	if flag {
		x = -x
	}

	for x > 0 {
		a, x = x%10, x/10
		list = append(list, a)
	}

	length = len(list)
	for i := 0; i < length; i++ {
		cnt := 0
		for cnt < length-i-1 {
			list[i] *= 10
			cnt += 1
		}
		t += int64(list[i])
	}

	if flag {
		if t > (1 << 31) {
			return 0

		}
		return -int(t-1) - 1
	}

	if t > ((1 << 31) - 1) {
		return 0
	}
	return int(t)
}

package main

func constructArr(a []int) []int {
	length := len(a)
	t := make([]int, length, length)
	t2 := make([]int, length, length)
	ans := make([]int, length, length)
	for i := range t {
		t[i] = 1
		t2[i] = 1
		ans[i] = 1
	}
	t[0] = a[0]
	t2[length-1] = a[length-1]
	for i := 1; i < length; i++ {
		t[i] = t[i-1] * a[i]
	}

	for i := 2; i <= length; i++ {
		t2[length-i] = t2[length-i+1]*a[length-i]
	}

	for i:= 0; i < length; i++ {
		if i > 0 {
			ans[i] *= t[i-1]
		}
		if i < length - 1{
			ans[i] *= t2[i+1]
		}
	}

	return ans
}

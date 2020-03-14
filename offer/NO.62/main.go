package main

func lastRemaining(n int, m int) int {
	ans := 0
	for i:= 2; i<=n;i++ {
		ans = (ans + m) % i
	}
	return ans
}

//func lastRemaining(n int, m int) int {
//	t := make([]int, n, n)
//	for i := 0; i < n; i++ {
//		t[i] = i
//	}
//	index := 0
//	for len(t) > 1 {
//		length := len(t)
//		index += m%length - 1 + length
//		index %= length
//		t = append(t[:index], t[index+1:]...)
//	}
//	return t[0]
//}

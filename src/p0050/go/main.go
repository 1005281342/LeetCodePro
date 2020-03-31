package main

// 快速幂
func myPow(x float64, n int) float64 {
	var ans float64 = 1
	flag := false
	if n < 0 {
		flag = true
	}

	for n > 0 {
		if n&1 > 0 {
			ans *= x
		}
		x *= x
		n /= 2
	}

	if flag {
		return 1 / ans
	}
	return ans
}

//func myPow(x float64, n int) float64 {
//	var ans float64 = 1
//	flag := false
//	if n < 0 {
//		flag = true
//		n = -n
//	}
//	for i := 1; i <= n; i++ {
//		ans *= x
//	}
//	if flag {
//		return 1 / ans
//	}
//	return ans
//}
//
//// 0.00001
//// 2147483647

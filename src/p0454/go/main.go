package main

func fourSumCount(A []int, B []int, C []int, D []int) int {

	var (
		length = len(A)
		mp     = make(map[int]int)
		cnt    int
	)

	if length == 0 {
		return 0
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			mp[A[i]+B[j]]++
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if v, ok := mp[-(C[i] + D[j])]; ok {
				cnt += v
			}
		}
	}

	return cnt
}

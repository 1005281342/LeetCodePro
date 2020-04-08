package main

func matrixBlockSum(mat [][]int, K int) [][]int {
	ans := make([][]int, len(mat), len(mat))
	for i := 0; i < len(mat); i++ {
		ans[i] = make([]int, len(mat[0]), len(mat[0]))
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[1]); j++ {
			cnt := 0
			for m := i - K; m < i+K+1; m++ {
				if m < 0 || m > len(mat) {
					continue
				}
				for n := j - K; n < j+K+1; n++ {
					if n < 0 || n > len(mat[0]) {
						continue
					}
				}
				cnt += mat[i][j]
			}
			ans[i][j] = cnt
		}
	}
	return ans
}

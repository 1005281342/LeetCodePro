package main

func rotate(matrix [][]int) {
	n := len(matrix)
	ans := make([][]int, n, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n, n)
		copy(ans[i], matrix[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = ans[i][j]
		}
	}
}

func main() {
	rotate([][]int{{0, 1}, {2, 3}})
}

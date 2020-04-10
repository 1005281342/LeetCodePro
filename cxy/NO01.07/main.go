package main

func rotate(matrix [][]int) {

	x := len(matrix)
	// 对角线交换
	for i := 0; i < x-1; i++ {
		for j := i + 1; j < x; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	//fmt.Println(matrix)
	// 对称交换
	for i := 0; i < x; i++ {
		for j := 0; j < x/2; j++ {
			matrix[i][j], matrix[i][x-1-j] = matrix[i][x-1-j], matrix[i][j]
		}
	}
	//fmt.Println(matrix)

	// 方法1
	//n := len(matrix)
	//ans := make([][]int, n, n)
	//for i := 0; i < n; i++ {
	//	ans[i] = make([]int, n, n)
	//	copy(ans[i], matrix[i])
	//}
	//
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n; j++ {
	//		matrix[j][n-i-1] = ans[i][j]
	//	}
	//}
}

func main() {
	rotate([][]int{{0, 1}, {2, 3}})
}

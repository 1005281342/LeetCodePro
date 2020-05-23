package main

func rotate(matrix [][]int) {
	//// 1. 2-5-8对称翻转
	//n := len(matrix)
	//for i := 0; i < n; i++ {
	//	for j := 0; i < (n >> 1); j++ {
	//		matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
	//	}
	//}/
	//
	//// 2. 以右对角线进行对称翻转(不好处理)

	// 1. 4-5-6对称翻转
	n := len(matrix)
	for j := 0; j < n; j++ {
		for i := 0; i < (n >> 1); i++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}

	// 2. 以左对角线进行对称翻转
	for i := 0; i < (n >> 1); i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}

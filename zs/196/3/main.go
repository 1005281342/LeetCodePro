package main

func numSubmat(mat [][]int) int {

	var (
		x    int
		y    int
		sums [][]int
		cnt  int
	)
	x = len(mat)
	if x == 0 {
		return 0
	}
	y = len(mat[0])
	if y == 0 {
		return 0
	}
	// 初始化
	sums = make([][]int, x+1, x+1)
	for i := 0; i < x+1; i++ {
		sums[i] = make([]int, y+1, y+1)
	}
	// 填充前缀和
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] - sums[i][j] + mat[i][j]
		}
	}

	// 查找前缀和
	// 扫描所有矩阵
	// 边长
	//var b = x
	//if y < x {
	//	b = y
	//}
	//for length:=1 ;length< b+1; length ++ {
	//	for i:= 0; i<x-length+1; i++ {
	//		for j:=0; j < y-length+1; j++ {
	//			if sums[i + length][j + length] - sums[i][j + length] - sums[i + length][j] + sums[i][j] == length * length {
	//				cnt++
	//			}
	//		}
	//	}
	//}
	for a := 1; a <= x; a++ {
		for b := 1; b <= y; b++ {
			for i := 0; i < x-a+1; i++ {
				for j := 0; j < y-b+1; j++ {
					if sums[i+a][j+b]-sums[i][j+b]-sums[i+a][j]+sums[i][j] == a*b {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}

package main

import (
	"fmt"
	"math"
)


func main() {

	var mat = [][]byte{
		[]byte{'0', '0', '1', '1'},
		[]byte{'0', '0', '1', '1'},
		[]byte{'0', '0', '1', '1'},
		[]byte{'0', '0', '1', '1'},
	}
	//var mat = [][]byte{
	//	[]byte{ '1', '1'},
	//	[]byte{ '1', '1'},
	//}
	fmt.Println(maximalSquare(mat))
}

// [["1","1"],["1","1"]]

func maximalSquare(matrix [][]byte) int {

	x := len(matrix)
	if x <= 0 {
		return 0
	}

	y := len(matrix[0])
	if y <= 0 {
		return 0
	}

	sum := make([][]int, x+1, x+1)
	for i := 0; i < x+1; i++ {
		sum[i] = make([]int, y+1, y+1)
	}

	k := make([][]int, x, x)
	for i := 0; i < x; i++ {
		k[i] = make([]int, y, y)
	}

	cnt := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if matrix[i][j] == '0' {
				k[i][j] = 0
			} else {
				k[i][j] = 1
				cnt += 1
			}
		}
	}

	if cnt == 0 {
		return 0
	}

	// 计算前缀和
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			sum[i+1][j+1] = k[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}

	//fmt.Println(sum)

	t := int(math.Sqrt(float64(cnt)))

	if y < t {
		t = y
	}

	if x < t {
		t = x
	}

	for t > 0 {
		ans := t * t

		for i := 0; i <= x-t; i++ {
			for j := 0; j <= y-t; j++ {
				if sum[t+i][t+j]+sum[i][j]-sum[i+t][j]-sum[i][j+t] == ans {
					return ans
				}
			}
		}
		t -= 1
	}

	return 1
}

package main

var MOVE [4][2]int = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func numRookCaptures(board [][]byte) int {
	x := len(board)
	if x <= 0 {
		return 0
	}
	y := len(board[0])

	// 找到'R'的位置
	i, j := 0, 0
	for ; i < x; i++ {
		f := false
		for j = 0; j < y; j++ {
			if board[i][j] == 'R' {
				f = true
				break
			}
		}
		if f {
			break
		}
	}

	var ans int
	for _, v := range MOVE {
		ti := i + v[0]
		tj := j + v[1]
		for 0 <= ti && ti < x && 0 <= tj && tj < y {
			if board[ti][tj] == 'p' {
				ans += 1
				break
			} else if board[ti][tj] != '.' {
				break
			}

			ti += v[0]
			tj += v[1]
		}
	}

	return ans
}

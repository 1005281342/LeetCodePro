package main

var MOVE_1 = [4][2]int{
	{-1, -1}, // 左上
	{-1, 0},  // 上
	{0, -1},  // 左
	{-1, 1},  // 右上
}
var MOVE_2 = [4][2]int{
	{1, 0},  // 下
	{0, 1},  // 右
	{1, 1},  // 右下
	{1, -1}, // 左下
}

func gameOfLife(board [][]int) {
	x := len(board)
	y := len(board[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			cnt := 0
			for _, a1 := range MOVE_1 {
				a1[0] += i
				a1[1] += j
				if 0 <= a1[0] && a1[0] < x && 0 <= a1[1] && a1[1] < y {
					if board[a1[0]][a1[1]] == 1 || board[a1[0]][a1[1]] == 2 {
						cnt += 1
					}
				}
			}

			for _, a1 := range MOVE_2 {
				a1[0] += i
				a1[1] += j
				if 0 <= a1[0] && a1[0] < x && 0 <= a1[1] && a1[1] < y {
					if board[a1[0]][a1[1]] == 1 {
						cnt += 1
					}
				}
			}

			if board[i][j] == 1 && (cnt < 2 || cnt > 3) {
				board[i][j] = 2 // 死亡 但之前是复活状态
			}
			if board[i][j] == 0 && cnt == 3 {
				board[i][j] = 3 // 复活 但之前是死亡状态
			}
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if board[i][j] >= 2 {
				board[i][j] -= 2
			}
		}
	}
}

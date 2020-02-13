package main

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	move := make(map[int][2]int, 4)
	move[0] = [2]int{0, 1}  // 右移
	move[1] = [2]int{1, 0}  // 下移
	move[2] = [2]int{0, -1} // 左移
	move[3] = [2]int{-1, 0} // 上移

	mark := make(map[[2]int]bool, len(matrix)*len(matrix[0]))
	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	ans = append(ans, matrix[0][0])
	mark[[2]int{0, 0}] = true

	x, y, now := 0, 0, 0
	for len(ans) < len(matrix)*len(matrix[0]) {
		a, b := move[now&3][0]+x, move[now&3][1]+y
		if (0 <= a && a < len(matrix)) && (0 <= b && b < len(matrix[0])) && !mark[[2]int{a, b}] {
			ans = append(ans, matrix[a][b])
			mark[[2]int{a, b}] = true
			x, y = a, b
		} else {
			now++
		}
	}
	return ans
}

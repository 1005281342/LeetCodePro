package main

func setZeroes(matrix [][]int) {
	x := len(matrix)
	if x <= 0 {
		return
	}
	y := len(matrix[0])

	t := make([][2]int, 0, 8)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if matrix[i][j] == 0 {
				t = append(t, [2]int{i, j})
			}
		}
	}

	for _, ij := range t {
		i, j := ij[0], ij[1]
		for m := 0; m < y; m++ {
			matrix[i][m] = 0
		}
		for m := 0; m < x; m++ {
			matrix[m][j] = 0
		}
	}

}

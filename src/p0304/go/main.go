package main

//
//func go() {
//
//}

type NumMatrix struct {
	DP   [][]int
	Flag bool
}

func Constructor(matrix [][]int) NumMatrix {

	if len(matrix) <= 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	dp := make([][]int, len(matrix)+1, len(matrix)+1)
	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			dp[i][j] = matrix[i-1][j-1] + dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1]
		}
	}
	//fmt.Println(dp)
	return NumMatrix{DP: dp, Flag: true}
}

func (this *NumMatrix) SumRegion(a int, b int, c int, d int) int {
	if !this.Flag {
		return 0
	}
	return this.DP[c+1][d+1] - this.DP[a][d+1] - this.DP[c+1][b] + this.DP[a][b]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

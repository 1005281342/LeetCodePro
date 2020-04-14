package main

func main() {
	findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
}

func findMaxForm(strs []string, m int, n int) int {

	length := len(strs)
	if length == 0 {
		return 0
	}

	var dp [][][]int
	dp = make([][][]int, length+1, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([][]int, m+1, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1, n+1)
		}
	}

	for i := 1; i <= length; i++ {
		mp := make(map[int32]int)
		for _, c := range strs[i-1] {
			mp[c] += 1
		}

		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if j >= mp['0'] && k >= mp['1'] {
					dp[i][j][k] = max(dp[i-1][j][k], 1+dp[i-1][j-mp['0']][k-mp['1']])
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[length][m][n]
}

//func findMaxForm(strs []string, m int, n int) int {
//
//	if len(strs) == 0 {
//		return 0
//	}
//
//	var dp [][]int
//	dp = make([][]int, m+1, m+1)
//	for i := 0; i <= m; i++ {
//		dp[i] = make([]int, n+1, n+1)
//	}
//	//fmt.Println(dp)
//
//	for _, str := range strs {
//		mp := make(map[int32]int)
//		for _, c := range str {
//			mp[c] += 1
//		}
//		for i := m; i >= mp['0']; i-- {
//			for j := n; j >= mp['1']; j-- {
//				dp[i][j] = max(1+dp[i-mp['0']][j-mp['1']], dp[i][j])
//				fmt.Println(dp)
//			}
//		}
//	}
//
//	return dp[m][n]
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

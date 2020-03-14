package main

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	dp := make([]int, length, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	maxLength := 1

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	//fmt.Println(dp)
	return maxLength
}

func main() {
	lengthOfLIS([]int{10,9,2,5,3,7,101,18})
}
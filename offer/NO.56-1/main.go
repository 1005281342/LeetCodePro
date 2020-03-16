package main

import "fmt"

func singleNumbers(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	low := xor ^ (xor-1)&xor
	fmt.Println(low)
	ans := make([]int, 2, 2)
	for _, v := range nums {
		t := ^v & low
		if t != 0 {
			ans[0] ^= v

		} else {
			ans[1] ^= v
		}
	}
	return ans
}

func main() {
	ans := singleNumbers([]int{4, 4, 2, 6})
	fmt.Println(ans)
}

/*
class Solution:
    def singleNumbers(self, nums: List[int]) -> List[int]:
        xor, ans = 0, [0, 0]
        for num in nums:
            xor ^= num
        low = xor ^ (xor - 1) & xor
        for num in nums:
            ans[not num & low] ^= num
        return ans
*/

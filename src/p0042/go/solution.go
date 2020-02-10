package main

func trap(height []int) int {
	if len(height) <= 0 {
		return 0
	}

	ans := 0
	mh := 0 // 找出最大的数
	for _, v := range height {
		if v > mh {
			mh = v
		}
	}

	for i := 1; i <= mh; i++ {
		left := 0
		right := len(height) - 1
		lok, rok := false, false
		for left < right && (!lok || !rok) {
			if height[left] < i {
				left += 1
			} else {
				lok = true
			}
			if height[right] < i {
				right -= 1
			} else {
				rok = true
			}
		}
		if left < right {
			for j := left + 1; j < right; j++ {
				if height[j] < i {
					ans += 1
				}
			}
		}
	}
	return ans
}

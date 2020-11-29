package main

func removeElement(nums []int, val int) int {
	var length int = len(nums)
	var (
		left  int
		right int = length - 1
		cnt   int
	)

	for right >= 0 {
		if nums[right] != val {
			break
		}
		cnt += 1
		right -= 1
	}

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right -= 1
			cnt += 1
			continue
		}
		left += 1
	}

	return length - cnt
}

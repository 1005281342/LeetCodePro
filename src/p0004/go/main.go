package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var (
		length int
		i      int
		j      int
		flag   bool
		index  int
		cnt    int
	)

	length = len(nums1) + len(nums2)
	flag = length&1 == 0
	index = (length + 1) >> 1
	if flag {
		index += 1
	}
	cnt = 0
	var x, y float64

	for cnt < index {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				y, x = float64(nums1[i]), y
				i++
			} else {
				y, x = float64(nums2[j]), y
				j++
			}
			cnt++
			continue
		}

		if i < len(nums1) {
			y, x = float64(nums1[i]), y
			i++
			cnt++
			continue
		}
		if j < len(nums2) {
			y, x = float64(nums2[j]), y
			j++
			cnt++
			continue
		}
		cnt++
	}

	if flag {
		return (x + y) / 2
	} else {
		return y
	}
}

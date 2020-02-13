package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	cur := head
	length := 0
	for cur != nil {
		length += 1
		cur = cur.Next
	}

	ans := make([]int, 0, length+10)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}

	left := 0
	right := len(ans) - 1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left += 1
		right -= 1
	}

	return ans
}

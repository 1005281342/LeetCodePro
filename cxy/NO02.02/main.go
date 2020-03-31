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

func kthToLast(head *ListNode, k int) int {
	cur := head
	for cur != nil {
		cur = cur.Next
		k -= 1
		if k < 0 {
			head = head.Next
		}
	}
	return head.Val
}

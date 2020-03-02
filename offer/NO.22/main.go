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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	a := head
	for k > 0 {
		a = a.Next
		k -= 1
	}

	for a != nil {
		a = a.Next
		head = head.Next
	}

	return head
}

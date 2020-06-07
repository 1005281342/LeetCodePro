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

func swapPairs(head *ListNode) *ListNode {
	a := head
	for a != nil && a.Next != nil {
		t := a.Next.Next
		a, a.Next = a.Next, a
		a = t
	}
	return head
}

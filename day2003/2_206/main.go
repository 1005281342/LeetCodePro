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

// 1, 2, 3, 4, 5, nil
// nil <- 1 <- 2 <- 3 <- 4 <- 5
func reverseList(head *ListNode) *ListNode {

	a := head
	var b *ListNode
	for a != nil {
		t := a.Next
		a.Next = b
		b = a
		a = t
	}
	return a
}

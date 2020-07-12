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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var temp *ListNode = head
	var k *ListNode = &ListNode{-1, head}
	var t *ListNode = k

	for n > 0 && temp != nil {
		temp = temp.Next
		n -= 1
	}

	for temp != nil {
		temp = temp.Next
		t = t.Next
	}

	if t.Next != nil {
		t.Next = t.Next.Next
		return k.Next
	}
	return nil
}

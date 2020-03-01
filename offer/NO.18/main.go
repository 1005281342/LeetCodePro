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

func deleteNode(head *ListNode, val int) *ListNode {

	if head.Val == val {
		return head.Next
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == val {
			if cur.Next.Next != nil {
				cur.Next = cur.Next.Next
			} else {
				cur.Next = nil
			}
		}
		cur = cur.Next
	}

	return head
}

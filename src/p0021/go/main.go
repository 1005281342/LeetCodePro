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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t := l1.Next
			l1.Next = nil
			cur.Next = l1
			cur = cur.Next
			l1 = t
		} else {
			t := l2.Next
			l2.Next = nil
			cur.Next = l2
			cur = cur.Next
			l2 = t
		}
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

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

	if head == nil {
		return nil
	}

	var (
		vn  *ListNode = &ListNode{Next: head}
		p   *ListNode = vn
		cur           = p.Next
	)

	for cur != nil && cur.Next != nil {
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p.Next.Next = cur

		p = cur
		cur = cur.Next
	}
	return vn.Next
}

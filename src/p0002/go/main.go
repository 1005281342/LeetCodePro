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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		t1 = l1
		t2 = l2
	)

	for t1 != nil && t2 != nil {
		t1 = t1.Next
		t2 = t2.Next
	}

	if t2 != nil {
		l1, l2 = l2, l1
	}

	t1, t2 = l1, l2

	for t1 != nil && t2 != nil {
		t1.Val += t2.Val
		t1 = t1.Next
		t2 = t2.Next
	}

	t1 = l1
	for t1 != nil && t1.Next != nil {
		if t1.Val >= 10 {
			t1.Next.Val += 1
			t1.Val -= 10
		}
		t1 = t1.Next
	}

	if t1 != nil && t1.Val >= 10 {
		t1.Val -= 10
		t1.Next = &ListNode{Val: 1}
	}

	return l1
}

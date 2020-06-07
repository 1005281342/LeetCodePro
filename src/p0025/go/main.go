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

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail, i := head, 0
	for ; tail != nil && i < k; i++ {
		tail = tail.Next
	}
	if i == k {
		tail = reverseKGroup(tail, k)
		for j := 0; j < k; j++ {
			tail, head.Next, head = head, tail, head.Next
		}
		head = tail
	}
	return head
}

//func reverseNodeList(node *ListNode) *ListNode {
//
//	// 1 -> 2 -> 3 -> nil
//	// nil <- 1 <- 2 <- 3
//	a := node
//	var b *ListNode
//	for a != nil {
//		t := a.Next // t = 2
//		a.Next = b  // a.next = nil
//		b = a       // b = 1
//		a = t       // a = t = 2  next
//	}
//
//	return node
//}

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

	var a = head
	for k > 0 {
		a = a.Next
		k -= 1
	}

	// 需要移除的节点个数和a当前地址到链表结尾节点数相等
	for a != nil {
		a = a.Next
		head = head.Next
	}

	return head
}

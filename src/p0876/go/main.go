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

func middleNode(head *ListNode) *ListNode {
	length := getLength(head)
	if length == 1 {
		return head
	}
	midIndex := length >> 1
	a := head
	cnt := 0
	for a != nil {
		a = a.Next
		cnt += 1
		if cnt == midIndex {
			break
		}
	}
	return a
}

func getLength(node *ListNode) int {
	cnt := 0
	for node != nil {
		cnt += 1
		node = node.Next
	}
	return cnt
}

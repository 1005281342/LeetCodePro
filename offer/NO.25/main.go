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

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var ans *ListNode
	if l2.Val > l1.Val {
		ans = &ListNode{
			Val:  l1.Val,
			Next: nil,
		}
		l1 = l1.Next
	} else {
		ans = &ListNode{
			Val:  l2.Val,
			Next: nil,
		}
		l2 = l2.Next
	}

	pre := ans
	for l1 != nil && l2 != nil {

		var a *ListNode
		if l2.Val > l1.Val {
			a = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			a = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		}
		pre.Next = a
		pre = pre.Next
	}

	if l1 != nil {
		pre.Next = l1
	}

	if l2 != nil {
		pre.Next = l2
	}

	return ans
}

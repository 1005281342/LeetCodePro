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
	var d1, d2 int
	a, b := l1, l2
	for a != nil {
		d1++
		a = a.Next
	}
	for b != nil {
		d2++
		b = b.Next
	}
	res, c := addTwoNumbers2(l1, l2, d1, d2)
	if c > 0 {
		return &ListNode{Val: c, Next: res}
	}
	return res
}
func addTwoNumbers2(l1, l2 *ListNode, d1, d2 int) (res *ListNode, c int) {
	if l1 != nil && l2 != nil {
		if l1.Next == nil && l2.Next == nil {
			n := l1.Val + l2.Val
			c = n / 10
			res = &ListNode{Val: n % 10, Next: nil}
			return
		}
	}
	var a *ListNode
	var b, n int
	if d1 > d2 {
		a, b = addTwoNumbers2(l1.Next, l2, d1-1, d2)
		n = l1.Val + b
	} else if d1 < d2 {
		a, b = addTwoNumbers2(l1, l2.Next, d1, d2-1)
		n = l2.Val + b
	} else {
		a, b = addTwoNumbers2(l1.Next, l2.Next, d1-1, d2-1)
		n = b + l1.Val + l2.Val
	}
	res = &ListNode{Val: n % 10, Next: a}
	c = n / 10
	return
}

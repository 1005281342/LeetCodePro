package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var mp map[int][]int

func helper(node *TreeNode, i int) {
	if node == nil {
		return
	}
	v, ok := mp[i]
	if ok {
		mp[i] = append(v, node.Val)
	} else {
		mp[i] = []int{node.Val}
	}
	helper(node.Left, i+1)
	helper(node.Right, i+1)
}

func list2ListNode(a []int) *ListNode {
	head := &ListNode{Val: a[0], Next: nil}
	cur := head
	for i := 1; i < len(a); i++ {
		cur.Next = &ListNode{Val: a[i]}
		cur = cur.Next
	}
	return head
}

func listOfDepth(tree *TreeNode) []*ListNode {

	mp = make(map[int][]int)
	helper(tree, 0)
	ans := make([]*ListNode, len(mp), len(mp))
	for k, v := range mp {
		ans[k] = list2ListNode(v)
	}
	return ans
}

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type H struct {
	T *TreeNode
	I int
}

var mp map[int]bool

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	mp = make(map[int]bool)
	mp[1] = true
	dq := make([]H, 0, 16)
	dq = append(dq, H{T: root, I: 1})
	ans := 1
	for len(dq) > 0 {
		h := dq[0]
		ans = h.I
		dq = dq[1:]
		if h.T.Left != nil {
			dq = append(dq, H{T:h.T.Left, I:h.I+1})
		}
		if h.T.Right != nil {
			dq = append(dq, H{T:h.T.Right, I:h.I+1})
		}
	}
	return ans
}

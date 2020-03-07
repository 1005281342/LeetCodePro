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

func levelOrder(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	ans := make([]int, 0, 1000)
	dq := make([]*TreeNode, 0, 1000)
	dq = append(dq, root)
	for len(dq) > 0 {
		node := dq[0]
		dq = dq[1:]
		ans = append(ans, node.Val)
		if node.Left != nil {
			dq = append(dq, node.Left)
		}
		if node.Right != nil {
			dq = append(dq, node.Right)
		}
	}
	return ans
}

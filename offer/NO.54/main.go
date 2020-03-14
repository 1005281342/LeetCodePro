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

var gCnt = 0
var ans = 0

func helper(node *TreeNode) {
	if node == nil || gCnt < 1 {
		return
	}

	if node.Right != nil {
		helper(node.Right)
	}

	if gCnt == 1 {
		ans = node.Val
		gCnt -= 1
		return
	}
	gCnt -= 1

	if node.Left != nil {
		helper(node.Left)
	}

}

func kthLargest(root *TreeNode, k int) int {
	gCnt = k
	helper(root)
	var t int
	t, ans = ans, t
	return t
}

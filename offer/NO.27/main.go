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

func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil {

		root.Right, root.Left = root.Left, root.Right

		mirrorTree(root.Left)
		mirrorTree(root.Right)
	}
	return root
}

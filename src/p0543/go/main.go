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

var ans = 1

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	L := helper(node.Left)
	R := helper(node.Right)
	t := L + R + 1
	if t > ans {
		ans = t
	}
	if L >= R {
		return L + 1
	} else {
		return R + 1
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	helper(root)
	t := ans - 1
	ans = 1
	return t
}

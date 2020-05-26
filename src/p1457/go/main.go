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

var t [9]int
var ans int

func pseudoPalindromicPaths(root *TreeNode) int {
	ans = 0
	t = [9]int{0, }
	helper(root)
	return ans
}

func helper(root *TreeNode)  {
	if root == nil {
		return
	}

	t[root.Val-1] ++

	if root.Left == nil && root.Right == nil {
		ans += check()
	}
	helper(root.Left)
	helper(root.Right)
	t[root.Val-1] --
}

func check() int {
	var cnt int
	for i := 0; i < 9; i++ {
		if t[i]&1 != 0 {
			cnt += 1
		}
	}
	if cnt > 1 {
		return 0
	}
	return 1
}

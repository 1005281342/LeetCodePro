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

func sumRootToLeaf(root *TreeNode) int {

	return dfs(root)
}

// 广度优先
func bfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var (
		dq  []*TreeNode
		ans int
	)
	dq = append(dq, node)

	for len(dq) > 0 {
		node = dq[0]
		dq = dq[1:]

		if node.Left != nil {
			node.Left.Val += node.Val * 2
			dq = append(dq, node.Left)
		}

		if node.Right != nil {
			node.Right.Val += node.Val * 2
			dq = append(dq, node.Right)
		}

		if node.Left == nil && node.Right == nil {
			ans += node.Val
		}
	}

	return ans
}

// 深度优先
func dfs(node *TreeNode) int {

	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		// 计算的值
		return node.Val
	}
	if node.Left != nil {
		node.Left.Val += 2 * node.Val
	}

	if node.Right != nil {
		node.Right.Val += 2 * node.Val
	}
	return dfs(node.Left) + dfs(node.Right)
}

package main

func main() {

}

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

func ok(s *TreeNode, t *TreeNode) bool {

	if s == nil && t == nil {
		return true
	}

	if (s == nil && t != nil) || (s != nil && t == nil) {
		return false
	}

	if s.Val != t.Val {
		return false
	}
	return ok(s.Left, t.Left) && ok(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {

	dq := make([]*TreeNode, 0, 16)
	dq = append(dq, s)
	for len(dq) > 0 {
		node := dq[0]
		dq = dq[1:]
		if ok(node, t) {
			return true
		}
		if node.Left != nil {
			dq = append(dq, node.Left)
		}
		if node.Right != nil {
			dq = append(dq, node.Right)
		}
	}
	return false
}

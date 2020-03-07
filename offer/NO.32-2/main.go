package main

import (
	"sort"
)

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

var t map[int][]int

func helper(node *TreeNode, i int) {
	if node == nil {
		return
	}
	_, ok := t[i]
	if !ok {
		t[i] = []int{node.Val}
	} else {
		t[i] = append(t[i], node.Val)
	}
	helper(node.Left, i+1)
	helper(node.Right, i+1)
}

func levelOrder(root *TreeNode) [][]int {
	t = make(map[int][]int)
	helper(root, 1)
	keys := make([]int, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	ans := make([][]int, 0, len(keys))
	for _, k := range keys {
		ans = append(ans, t[k])
	}
	return ans
}

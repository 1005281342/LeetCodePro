package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 0 {
		return true
	}
	return helper(postorder)
}

func helper(a []int) bool {
	length := len(a)
	if length <= 0 {
		return true
	}
	root := a[length-1] // 根结点
	i := -1

	// 左子树
	for i < length-1 {
		i += 1
		if a[i] > root { // 如果这个值大于根的值，说明若是BST该值应在右边
			break
		}
	}
	// 右子树
	for j := i; j < length-1; j++ {
		if a[j] < root { // 若右子树存在值小于根节点， 则不是BST
			return false
		}
	}
	return helper(a[0:i]) && helper(a[i:length-1])
}

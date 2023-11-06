package levelPrint

import "dfs_leetcode/utils"

func rightSideView(root *utils.TreeNode) []int {
	store := new([][]int)

	if root == nil {
		return []int{}
	}

	levelTraversal(store, root, 0)

	res := []int{}
	for i := 0; i < len(*store); i++ {
		length := len((*store)[i])
		res = append(res, (*store)[i][length-1])
	}

	return res
}

func levelTraversal(store *[][]int, root *utils.TreeNode, level int) {
	if root == nil {
		return
	}

	if (*store)[level] == nil {
		(*store)[level] = make([]int, 0)
		(*store)[level] = append((*store)[level], root.Val)
	} else {
		(*store)[level] = append((*store)[level], root.Val)
	}

	levelTraversal(store, root.Left, level+1)
	levelTraversal(store, root.Right, level+1)
}

package levelPrint

import "dfs_leetcode/utils"

func rightSideView(root *utils.TreeNode) []int {
	store := new([][]int)

	if root == nil {
		return []int{}
	}

	levelTraversal(store, root, 1)

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

	if (*store) == nil || len(*store) < level {
		*store = append(*store, []int{root.Val}) // new level, no need to append, directly add new array with val in it.
	} else {
		(*store)[level-1] = append((*store)[level-1], root.Val)
	}

	levelTraversal(store, root.Left, level+1)
	levelTraversal(store, root.Right, level+1)
}

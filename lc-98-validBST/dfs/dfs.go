package dfs

import (
	"dfs_leetcode/utils"
	"math"
)

func isValidBST(root *utils.TreeNode) bool {
	return dfs(math.MinInt, math.MaxInt, root)
}

func dfs(min, max int, node *utils.TreeNode) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return dfs(min, node.Val, node.Left) && dfs(node.Val, max, node.Right)
}

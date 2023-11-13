package dfs

import "dfs_leetcode/utils"

func DiameterOfBinaryTree(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	rootLen := dfs(root.Left) + dfs(root.Right)
	leftLen := DiameterOfBinaryTree(root.Left)
	rightLen := DiameterOfBinaryTree(root.Right)

	return maxThree(rootLen, leftLen, rightLen)
}

func dfs(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	leftLen := 1 + dfs(root.Left)
	rightLen := 1 + dfs(root.Right)

	return maxTwo(leftLen, rightLen)
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxThree(a, b, c int) int {
	arr := []int{a, b, c}
	maxVal := 0
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

var DP map[*utils.TreeNode]int // global map
func DiameterOfBinaryTreeDP(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	dp := make(map[*utils.TreeNode]int)

	leftLen := 0
	if _, ok := dp[root.Left]; ok {
		leftLen = 1 + dp[root.Left]
	} else {
		leftLen = 1 + dfsDP(root.Left, dp)
	}

	rightLen := 0
	if _, ok := dp[root.Right]; ok {
		rightLen = 1 + dp[root.Right]
	} else {
		rightLen = 1 + dfsDP(root.Right, dp)
	}

	rootPath := leftLen + rightLen - 2
	leftNode := DiameterOfBinaryTreeDP(root.Left)
	rightNode := DiameterOfBinaryTreeDP(root.Right)

	return maxThree(rootPath, leftNode, rightNode)
}

func dfsDP(root *utils.TreeNode, dp map[*utils.TreeNode]int) int {
	if root == nil {
		return 0
	}

	leftLen := 0
	if _, ok := dp[root.Left]; ok {
		leftLen = 1 + dp[root.Left]
	} else {
		leftLen = 1 + dfsDP(root.Left, dp)
	}

	rightLen := 0
	if _, ok := dp[root.Right]; ok {
		rightLen = 1 + dp[root.Right]
	} else {
		rightLen = 1 + dfsDP(root.Right, dp)
	}

	dp[root] = maxTwo(leftLen, rightLen)
	return dp[root]
}

func diameterOfBinaryTree(root *utils.TreeNode) int {
	res := 0

	dfsIdeal(root, &res)

	return res
}

func dfsIdeal(root *utils.TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := dfsIdeal(root.Left, res)
	right := dfsIdeal(root.Right, res)

	*res = maxTwo(*res, left+right)

	return 1 + maxTwo(left, right)
}

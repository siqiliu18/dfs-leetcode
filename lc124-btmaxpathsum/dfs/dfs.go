package dfs

import "dfs_leetcode/utils"

// Note: sum of a path, not sum of subTree!!!
// [5,4,8,11,null,13,4,7,2,null,null,null,1], the path is 7 -> 11 -> 4 -> 5 -> 8 -> 13 = 48

var minVal int = -1000

func MaxpathSum(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	leftPathSum, rightPathSum, leftPathDFS, rightPathDFS := minVal, minVal, minVal, minVal
	bothPathSum, rootLeftSum, rootRightSum := root.Val, root.Val, root.Val
	if root.Left != nil {
		leftPathSum = maxSumOfARoot(root.Left)
		bothPathSum += leftPathSum
		rootLeftSum += leftPathSum
		leftPathDFS = MaxpathSum(root.Left)
	}

	if root.Right != nil {
		rightPathSum = maxSumOfARoot(root.Right)
		bothPathSum += rightPathSum
		rootRightSum += rightPathSum
		rightPathDFS = MaxpathSum(root.Right)
	}

	return findMax(root.Val, leftPathSum, rightPathSum, bothPathSum, rootLeftSum, rootRightSum, leftPathDFS, rightPathDFS)
}

func maxSumOfARoot(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	return findMaxOfTwo(root.Val, root.Val+findMaxOfTwo(maxSumOfARoot(root.Left), maxSumOfARoot(root.Right)))
}

func findMaxOfTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(rootVal, leftVal, rightVal, bothPathVal, rootLeftVal, rootRightVal, leftPath, rightPath int) int {
	arr := []int{}
	arr = append(arr, rootVal)
	arr = append(arr, leftVal)
	arr = append(arr, rightVal)
	arr = append(arr, bothPathVal)
	arr = append(arr, rootLeftVal)
	arr = append(arr, rootRightVal)
	arr = append(arr, leftPath)
	arr = append(arr, rightPath)

	maxVal := minVal

	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

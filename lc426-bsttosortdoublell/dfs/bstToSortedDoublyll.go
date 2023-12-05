package dfs

import (
	"dfs_leetcode/utils"
)

var Head = utils.TreeNode{}
var Pre = &Head
var Curr = Head.Right // nil

func TreeToDoublyList(root *utils.TreeNode) *utils.TreeNode {
	inorder(root)
	Pre.Right = Head.Right
	Head.Right.Left = Pre
	return Head.Right
}

func inorder(root *utils.TreeNode) {
	if root == nil {
		return
	} else {
		inorder(root.Left)
		Curr = &utils.TreeNode{
			Val:  root.Val,
			Left: Pre,
		}
		Pre.Right = Curr
		Pre = Curr
		Curr = Curr.Right
		inorder(root.Right)
	}
}

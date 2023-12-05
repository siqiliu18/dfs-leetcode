package dfs

import (
	"dfs_leetcode/utils"
)

var Head = utils.TreeNode{}
var Pre = &Head
var Curr = Head.Right // nil

func TreeToDoublyList(root *utils.TreeNode) *utils.TreeNode {
	head := utils.TreeNode{}
	pre := &head
	curr := head.Right // nil

	var dfs func(*utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		curr = &utils.TreeNode{
			Val:  node.Val,
			Left: pre,
		}
		pre.Right = curr
		pre = curr
		curr = curr.Right
		dfs(node.Right)
	}

	dfs(root)
	pre.Right = head.Right
	head.Right.Left = pre
	return head.Right
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

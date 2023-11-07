package dfs

import "dfs_leetcode/utils"

/*
case 1: if root != p != q and p found on left and q found or right, return root
case 2: if root = p and q found on the right, return p or root
case 3: if q found on the left and root = q, return q or root
*/

func LowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if found(root.Left, p) && found(root.Right, q) {
		return root
	}

	leftRes := LowestCommonAncestor(root.Left, p, q)
	if leftRes != nil {
		return leftRes
	}

	rightRes := LowestCommonAncestor(root.Right, p, q)
	if rightRes != nil {
		return rightRes
	}

	return nil
}

func found(root, node *utils.TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == node.Val {
		return true
	}

	return found(root.Left, node) || found(root.Right, node)
}

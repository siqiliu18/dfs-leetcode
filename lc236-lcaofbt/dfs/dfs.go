package dfs

import (
	"dfs_leetcode/utils"
)

/*
case 1: if root != p != q and p found on left and q found or right, return root
case 2: if root = p and q found on the right, return p or root
case 3: if q found on the left and root = q, return q or root
*/

func LowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if (found(root.Left, p) && found(root.Right, q)) || (found(root.Left, q) && found(root.Right, p)) {
		return root
	}

	if (root.Val == p.Val && found(root.Right, q)) || (root.Val == p.Val && found(root.Left, q)) {
		return root
	}

	if (found(root.Left, p) && root.Val == q.Val) || (found(root.Right, p) && root.Val == q.Val) {
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

func nodeAddr(root, node *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == node.Val {
		return root
	}

	res := nodeAddr(root.Left, node)
	if res != nil {
		return res
	}

	return nodeAddr(root.Right, node)
}

// map children to parents
func LowestCommonAncestor2(root, p, q *utils.TreeNode) *utils.TreeNode {
	childToParents := make(map[*utils.TreeNode]*utils.TreeNode)
	fillMap(childToParents, root)
	childToParents[root] = nil

	pParents := make(map[*utils.TreeNode]int)
	getParents(p, childToParents, pParents)

	return findParent(q, childToParents, pParents)
}

func fillMap(childToParents map[*utils.TreeNode]*utils.TreeNode, root *utils.TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		childToParents[root.Left] = root
	}

	if root.Right != nil {
		childToParents[root.Right] = root
	}

	fillMap(childToParents, root.Left)
	fillMap(childToParents, root.Right)
}

func getParents(node *utils.TreeNode, childToParents map[*utils.TreeNode]*utils.TreeNode, parents map[*utils.TreeNode]int) {
	if node == nil {
		return
	}

	parents[node] = 0

	nextNode := childToParents[node]

	getParents(nextNode, childToParents, parents)
}

func findParent(node *utils.TreeNode, childToParents map[*utils.TreeNode]*utils.TreeNode, parents map[*utils.TreeNode]int) *utils.TreeNode {
	if _, found := parents[node]; found {
		return node
	}

	return findParent(childToParents[node], childToParents, parents)
}

// faster
func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

package dfs

import (
	"dfs_leetcode/utils"
	"fmt"
	"sort"
)

// idea from https://leetcode.com/problems/binary-tree-vertical-order-traversal/solutions/1852164/golang-solution-with-explanation-and-images/
func VerticalOrder(root *utils.TreeNode) [][]int {
	vMap := make(map[int][]int)

	// inorder(0, root, vMap)
	bfs(root, vMap)

	fmt.Println(vMap)

	sortKey := []int{}

	for k := range vMap {
		sortKey = append(sortKey, k)
	}

	sort.SliceStable(sortKey, func(i, j int) bool {
		return sortKey[i] < sortKey[j]
	})

	fmt.Println(sortKey)

	res := new([][]int)
	for _, k := range sortKey {
		*res = append(*res, vMap[k])
	}

	return *res
}

func inorder(col int, root *utils.TreeNode, vMap map[int][]int) {
	if root == nil {
		return
	}

	inorder(col-1, root.Left, vMap)

	vMap[col] = append(vMap[col], root.Val)

	inorder(col+1, root.Right, vMap)
}

type key struct {
	node *utils.TreeNode
	col  int // the column of the node
}

func bfs(root *utils.TreeNode, vMap map[int][]int) {
	queue := []key{{root, 0}}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.node == nil {
			continue
		}

		vMap[cur.col] = append(vMap[cur.col], cur.node.Val)

		queue = append(queue, key{cur.node.Left, cur.col - 1}, key{cur.node.Right, cur.col + 1})
	}
}

package levelPrint

import "dfs_leetcode/utils"

func RightSideView(root *utils.TreeNode) []int {
	// this article talks about diff of new and make https://stackoverflow.com/questions/9320862/why-would-i-make-or-new
	/*	Quote:

		*The reason is that slice, map and chan are data structures. They need to be initialized, otherwise they won't be usable.

		* new(T) - Allocates memory, and sets it to the zero value for type T....that is 0 for int, "" for string and nil for referenced types (slice, map, chan)
			Note that referenced types are just pointers to some underlying data structures, which won't be created by new(T)
			Example: in case of slice, the underlying array won't be created, thus new([]int) returns a pointer to nothing

		* make(T) - Allocates memory for referenced data types (slice, map, chan), plus initializes their underlying data structures
	*/
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

// I am traversing all nodes, so the big O should just be O(n)
// https://stackoverflow.com/questions/4547012/complexities-of-binary-tree-traversals
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

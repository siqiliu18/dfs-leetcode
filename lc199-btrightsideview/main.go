package main

import (
	"dfs_leetcode/utils"
	"fmt"

	dfs "dfs_leetcode/lc199-btrightsideview/levelPrint"
)

func main() {
	input := "[1,2,3,4]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	res := dfs.RightSideView(root)
	fmt.Println(res)
}

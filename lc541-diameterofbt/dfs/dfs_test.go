package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDFS(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := dfs(root)
		So(res, ShouldEqual, 3)
	})
	Convey("2", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := dfs(root.Left)
		So(res, ShouldEqual, 2)
	})
	Convey("3", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := dfs(root.Right)
		So(res, ShouldEqual, 1)
	})
}

func TestDiameterOfBinaryTree(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTree(root)
		So(res, ShouldEqual, 3)
	})
	Convey("2", t, func() {
		input := "[1,2,3,4,5,null,null,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTree(root)
		So(res, ShouldEqual, 4)
	})
	Convey("3", t, func() {
		input := "[1,2,3,4,5,null,null,null,null,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTree(root)
		So(res, ShouldEqual, 4)
	})
	Convey("4", t, func() {
		input := "[1,2,null,4,5,null,8,null,7,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTree(root)
		So(res, ShouldEqual, 5)
	})
}

func TestDFS_DP(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		dp := make(map[*utils.TreeNode]int)
		res := dfsDP(root, dp)
		So(res, ShouldEqual, 3)
	})
	Convey("2", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		dp := make(map[*utils.TreeNode]int)
		res := dfsDP(root.Left, dp)
		So(res, ShouldEqual, 2)
	})
	Convey("3", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		dp := make(map[*utils.TreeNode]int)
		res := dfsDP(root.Right, dp)
		So(res, ShouldEqual, 1)
	})
	Convey("4", t, func() {
		input := "[1,2,null,4,5,null,8,null,7,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		dp := make(map[*utils.TreeNode]int)
		res := dfsDP(root, dp)
		So(res, ShouldEqual, 5)
		So(dp[root], ShouldEqual, 5)
		So(dp[root.Left], ShouldEqual, 4)
		So(dp[root.Right], ShouldEqual, 0)
	})
	Convey("5", t, func() {
		input := "[1]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		dp := make(map[*utils.TreeNode]int)
		res := dfsDP(root, dp)
		So(res, ShouldEqual, 1)
		So(dp[root], ShouldEqual, 1)
	})
}

func TestDiameterOfBinaryTreeDP(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTreeDP(root)
		So(res, ShouldEqual, 3)
	})
	Convey("2", t, func() {
		input := "[1,2,3,4,5,null,null,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTreeDP(root)
		So(res, ShouldEqual, 4)
	})
	Convey("3", t, func() {
		input := "[1,2,3,4,5,null,null,null,null,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTreeDP(root)
		So(res, ShouldEqual, 4)
	})
	Convey("4", t, func() {
		input := "[1,2,null,4,5,null,8,null,7,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTreeDP(root)
		So(res, ShouldEqual, 5)
	})
	Convey("5", t, func() {
		input := "[1]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := DiameterOfBinaryTreeDP(root)
		So(res, ShouldEqual, 0)
	})
}

func TestDiameterOfBinaryTreeIdeal(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,4,5]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := diameterOfBinaryTree(root)
		So(res, ShouldEqual, 3)
	})
	Convey("2", t, func() {
		input := "[1,2,3,4,5,null,null,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := diameterOfBinaryTree(root)
		So(res, ShouldEqual, 4)
	})
	Convey("3", t, func() {
		input := "[1,2,3,4,5,null,null,null,null,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := diameterOfBinaryTree(root)
		So(res, ShouldEqual, 4)
	})
	Convey("4", t, func() {
		input := "[1,2,null,4,5,null,8,null,7,6]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := diameterOfBinaryTree(root)
		So(res, ShouldEqual, 5)
	})
	Convey("5", t, func() {
		input := "[1]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := diameterOfBinaryTree(root)
		So(res, ShouldEqual, 0)
	})
}

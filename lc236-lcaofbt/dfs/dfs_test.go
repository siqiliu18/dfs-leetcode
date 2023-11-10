package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFound(t *testing.T) {
	input := "[3,5,1,6,2,0,8,null,null,7,4]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		p := utils.TreeNode{
			Val: 0,
		}
		res := found(root, &p)
		So(res, ShouldBeTrue)
	})
	Convey("2", t, func() {
		p := utils.TreeNode{
			Val: 8,
		}
		res := found(root, &p)
		So(res, ShouldBeTrue)
	})
}

func TestRightSideView1(t *testing.T) {
	Convey("1", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		p := utils.TreeNode{
			Val: 0,
		}
		q := utils.TreeNode{
			Val: 8,
		}
		res := LowestCommonAncestor(root, &p, &q)
		So(res.Val, ShouldResemble, 1)
	})
}

func TestRightSideView2(t *testing.T) {
	Convey("2", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		p := utils.TreeNode{
			Val: 4,
		}
		q := utils.TreeNode{
			Val: 8,
		}
		res := LowestCommonAncestor(root, &p, &q)
		So(res.Val, ShouldResemble, 3)
	})
}

func TestRightSideView3(t *testing.T) {
	Convey("3", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		p := utils.TreeNode{
			Val: 5,
		}
		q := utils.TreeNode{
			Val: 4,
		}
		res := LowestCommonAncestor(root, &p, &q)
		So(res.Val, ShouldResemble, 5)
	})
}

func TestRightSideView4(t *testing.T) {
	Convey("4", t, func() {
		input := "[1,2]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		p := utils.TreeNode{
			Val: 1,
		}
		q := utils.TreeNode{
			Val: 2,
		}
		res := LowestCommonAncestor(root, &p, &q)
		So(res.Val, ShouldResemble, 1)
	})
}

func TestLowestCommonAncestor2(t *testing.T) {
	Convey("1", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		p := utils.TreeNode{
			Val: 6,
		}
		pAddr := nodeAddr(root, &p)
		q := utils.TreeNode{
			Val: 8,
		}
		qAddr := nodeAddr(root, &q)
		res := LowestCommonAncestor2(root, pAddr, qAddr)
		So(res.Val, ShouldResemble, 3)
	})
}

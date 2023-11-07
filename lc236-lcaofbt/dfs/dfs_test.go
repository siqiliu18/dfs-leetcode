package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFound(t *testing.T) {
	input := "[3,5,1,6,2,0,8,null,null,7,4]"
	root := utils.BuildTreeFromStringArray(input)
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

func TestRightSideView(t *testing.T) {
	Convey("1", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArray(input)
		p := utils.TreeNode{
			Val: 0,
		}
		q := utils.TreeNode{
			Val: 8,
		}
		res := LowestCommonAncestor(root, &p, &q)
		So(res.Val, ShouldResemble, 1)
	})
	Convey("2", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArray(input)
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

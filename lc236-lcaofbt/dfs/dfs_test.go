package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRightSideView(t *testing.T) {
	Convey("1", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"
		root := utils.BuildTreeFromStringArray(input)
		p := utils.TreeNode{
			Val: 6,
		}
		q := utils.TreeNode{
			Val: 4,
		}
		res := LowestCommonAncestor(root, &p, &q)
		So(res.Val, ShouldResemble, 5)
	})
}

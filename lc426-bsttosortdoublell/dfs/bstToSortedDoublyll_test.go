package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTreeToDoublyList1(t *testing.T) {
	input := "[4,2,5,1,3]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := TreeToDoublyList(root)
		So(res.Val, ShouldEqual, 1)
		So(res.Left.Val, ShouldEqual, 5)
	})
}

func TestTreeToDoublyList2(t *testing.T) {
	input := "[2,1,3]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := TreeToDoublyList(root)
		So(res.Val, ShouldEqual, 1)
		So(res.Left.Val, ShouldEqual, 3)
	})
}

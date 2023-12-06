package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTreeToDoublyList1(t *testing.T) {
	input := "[3,9,8,4,0,1,7,null,null,null,2,5]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := VerticalOrder(root)
		So(res, ShouldEqual, [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}})
	})
}

func TestTreeToDoublyList2(t *testing.T) {
	input := "[3,9,8,4,0,1,7]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := VerticalOrder(root)
		So(res, ShouldEqual, [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}})
	})
}

func TestTreeToDoublyList3(t *testing.T) {
	input := "[3,9,20,null,null,15,7]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := VerticalOrder(root)
		So(res, ShouldEqual, [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}})
	})
}

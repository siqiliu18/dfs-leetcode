package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxPathSum1(t *testing.T) {
	input := "[1,2,3]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 6)
	})
}

func TestMaxPathSum2(t *testing.T) {
	input := "[-10,9,20,null,null,15,7]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 42)
	})
}

func TestMaxPathSum3(t *testing.T) {
	input := "[-10,-9,-20,null,null,15,7]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 15)
	})
}

func TestMaxPathSum4(t *testing.T) {
	input := "[1,-2,3]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 4)
	})
}

func TestMaxPathSum6(t *testing.T) {
	input := "[-3]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, -3)
	})
}

func TestMaxPathSum7(t *testing.T) {
	input := "[-2,-1]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, -1)
	})
}

func TestMaxPathSum8(t *testing.T) {
	input := "[-2,-3,-1]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, -1)
	})
}

func TestMaxPathSum9(t *testing.T) {
	input := "[1,-2,-3,1,3,-2,null,-1]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 3)
	})
}

func TestMaxPathSum10(t *testing.T) {
	input := "[-1,-2,10,-6,null,-3,-6]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 10)
	})
}

func TestMaxPathSum5(t *testing.T) {
	input := "[5,4,8,11,null,13,4,7,2,null,null,null,1]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := MaxpathSum(root)
		So(res, ShouldEqual, 48)
	})
}

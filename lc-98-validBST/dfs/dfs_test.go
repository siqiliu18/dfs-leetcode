package dfs

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsValidBST1(t *testing.T) {
	input := "[2,1,3]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := isValidBST(root)
		So(res, ShouldBeTrue)
	})
}

func TestIsValidBST12(t *testing.T) {
	input := "[5,1,4,null,null,3,6]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := isValidBST(root)
		So(res, ShouldBeFalse)
	})
}

func TestIsValidBST13(t *testing.T) {
	input := "[5,1,8,null,null,6,9]"
	root := utils.BuildTreeFromStringArrayForLoop(input)
	Convey("1", t, func() {
		res := isValidBST(root)
		So(res, ShouldBeTrue)
	})
}

package levelPrint

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRightSideView(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,4]"
		root := utils.BuildTreeFromStringArrayForLoop(input)
		res := RightSideView(root)
		So(res, ShouldResemble, []int{1, 3, 4})
	})
}

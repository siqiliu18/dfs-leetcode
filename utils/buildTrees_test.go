package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildTreeFromStringArray(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,null,5,null,4]"

		res := BuildTreeFromStringArray(input)

		So(res.Left.Val, ShouldEqual, 2)
		So(res.Left.Left, ShouldBeNil)
		So(res.Right.Right.Val, ShouldEqual, 4)
	})
}

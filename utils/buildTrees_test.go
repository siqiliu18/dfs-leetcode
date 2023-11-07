package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildTreeFromStringArray1(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,null,5,null,4]"

		res := BuildTreeFromStringArray(input)

		So(res.Left.Val, ShouldEqual, 2)
		So(res.Left.Left, ShouldBeNil)
		So(res.Right.Right.Val, ShouldEqual, 4)
	})
}

func TestBuildTreeFromStringArray2(t *testing.T) {
	Convey("2", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"

		res := BuildTreeFromStringArray(input)

		So(res.Left.Val, ShouldEqual, 5)
		So(res.Left.Left.Left, ShouldBeNil)
		So(res.Left.Right.Left.Val, ShouldEqual, 7)
		So(res.Right.Right.Val, ShouldEqual, 8)
	})
}

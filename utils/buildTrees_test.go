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

func TestBuildTreeFromStringArray3(t *testing.T) {
	Convey("2", t, func() {
		input := "[5,4,8,11,null,13,4,7,2,null,null,null,1]"

		res := BuildTreeFromStringArray(input)

		So(res.Right.Left.Right, ShouldBeNil)
		So(res.Right.Right.Right.Val, ShouldEqual, 1)
	})
}

func TestBuildTreeFromStringArrayForLoop1(t *testing.T) {
	Convey("1", t, func() {
		input := "[5,4,8,11,null,13,4,7,2,null,null,null,1]"

		res := BuildTreeFromStringArrayForLoop(input)

		So(res.Right.Left.Right, ShouldBeNil)
		So(res.Right.Right.Right.Val, ShouldEqual, 1)
		So(res.Right.Right.Left, ShouldBeNil)
	})
}

func TestBuildTreeFromStringArrayForLoop2(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,null,5,null,4]"

		res := BuildTreeFromStringArrayForLoop(input)

		So(res.Left.Val, ShouldEqual, 2)
		So(res.Left.Left, ShouldBeNil)
		So(res.Right.Right.Val, ShouldEqual, 4)
	})
}

func TestBuildTreeFromStringArrayForLoop3(t *testing.T) {
	Convey("2", t, func() {
		input := "[3,5,1,6,2,0,8,null,null,7,4]"

		res := BuildTreeFromStringArrayForLoop(input)

		So(res.Left.Val, ShouldEqual, 5)
		So(res.Left.Left.Left, ShouldBeNil)
		So(res.Left.Right.Left.Val, ShouldEqual, 7)
		So(res.Right.Right.Val, ShouldEqual, 8)
	})
}

func TestBuildTreeFromStringArrayForLoop4(t *testing.T) {
	Convey("2", t, func() {
		input := "[-2,-1]"

		res := BuildTreeFromStringArrayForLoop(input)

		So(res.Left.Val, ShouldEqual, -1)
	})
}

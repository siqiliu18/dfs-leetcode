package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildTreeFromStringArray(t *testing.T) {
	Convey("1", t, func() {
		input := "[1,2,3,null,5,null,4]"

		res := BuildTreeFromStringArray(input)

		So(res, ShouldEqual, "1,2,3,null,5,null,4")
	})
}

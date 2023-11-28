package lc269aliendict

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExist1(t *testing.T) {
	Convey("1", t, func() {
		words := []string{"wrt", "wrf", "er", "ett", "rftt"}
		res := AlienOrderBFS(words)
		So(res, ShouldEqual, "wertf")
	})
}

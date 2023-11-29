package lc269aliendict

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExistBFS1(t *testing.T) {
	Convey("1", t, func() {
		words := []string{"wrt", "wrf", "er", "ett", "rftt"}
		res := AlienOrderBFS(words)
		So(res, ShouldEqual, "wertf")
	})
}

func TestExistBFS2(t *testing.T) {
	Convey("2", t, func() {
		words := []string{"zy", "zx"}
		res := AlienOrderBFS(words)
		So(res, ShouldEqual, "yxz")
	})
}

func TestExistBFS3(t *testing.T) {
	Convey("2", t, func() {
		words := []string{"ac", "ab", "b"}
		res := AlienOrderBFS(words)
		So(res, ShouldEqual, "acb")
	})
}

func TestExistDFS2(t *testing.T) {
	Convey("2", t, func() {
		words := []string{"wrt", "wrf", "er", "ett", "rftt"}
		res := AlienOrderPostOrderDFS(words)
		So(res, ShouldEqual, "wertf")
	})
}

func TestExistDFS3(t *testing.T) {
	Convey("3", t, func() {
		words := []string{"z", "x", "z"}
		res := AlienOrderPostOrderDFS(words)
		So(res, ShouldEqual, "")
	})
}

func TestExistDFS4(t *testing.T) {
	Convey("4", t, func() {
		words := []string{"zy", "zx"}
		res := AlienOrderPostOrderDFS(words)
		So(res, ShouldEqual, "yxz")
	})
}

func TestExistDFS5(t *testing.T) {
	Convey("5", t, func() {
		words := []string{"ac", "ab", "b"}
		res := AlienOrderPostOrderDFS(words)
		So(res, ShouldEqual, "acb")
	})
}

func TestExistDFS6(t *testing.T) {
	Convey("6", t, func() {
		words := []string{"z", "x", "a", "zb", "zx"}
		res := AlienOrderPostOrderDFS(words)
		So(res, ShouldEqual, "acb")
	})
}

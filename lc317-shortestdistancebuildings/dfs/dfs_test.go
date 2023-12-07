package dfs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortestDistance0(t *testing.T) {
	Convey("0", t, func() {
		input := [][]int{{1, 0}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, 1)
	})
}

func TestShortestDistance1(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 0, 2, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, 7)
	})
}

func TestShortestDistance2(t *testing.T) {
	Convey("2", t, func() {
		input := [][]int{{1, 0, 2, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, 10)
	})
}

func TestShortestDistance3(t *testing.T) {
	Convey("3", t, func() {
		input := [][]int{{0, 0, 1, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, 8)
	})
}

func TestShortestDistance31(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 0, 1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, 2)
	})
}

func TestShortestDistance4(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, -1)
	})
}

func TestShortestDistance5(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, -1)
	})
}

func TestShortestDistance6(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1}, {1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, -1)
	})
}

func TestShortestDistance7(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1}, {0}, {1}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, 2)
	})
}

func TestShortestDistance8(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 2, 0}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, -1)
	})
}

func TestShortestDistance9(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{0, 2, 1}, {1, 0, 2}, {0, 1, 0}}
		res := ShortestDistance(input)
		So(res, ShouldEqual, -1)
	})
}

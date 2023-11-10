package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildGraphsUsingQueue(t *testing.T) {
	Convey("1", t, func() {
		input := "[[2,4],[1,3],[2,4],[1,3]]"
		res := BuildGraphsBFS(input)
		fmt.Println(res)
		So(res.Neighbors[0].Val, ShouldEqual, 2)
		So(res.Neighbors[1].Neighbors[1].Val, ShouldEqual, 3)
	})
}

func TestBuildGraphsDFS(t *testing.T) {
	Convey("1", t, func() {
		input := "[[2,4],[1,3],[2,4],[1,3]]"
		res := BuildGraphsDFS(input)
		fmt.Println(res)
		So(res.Neighbors[0].Val, ShouldEqual, 2)
		So(res.Neighbors[1].Neighbors[1].Val, ShouldEqual, 3)
	})
}

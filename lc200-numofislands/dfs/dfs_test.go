package dfs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExist1(t *testing.T) {
	Convey("1", t, func() {
		grid2 := [][]byte{
			{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
			{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
			{byte('0'), byte('0'), byte('1'), byte('0'), byte('0')},
			{byte('0'), byte('0'), byte('0'), byte('1'), byte('1')},
		}

		res := NumIslands(grid2)

		So(res, ShouldEqual, 3)
	})
}

package dfs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExist1(t *testing.T) {
	Convey("1", t, func() {
		testBoard := [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('C'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}
		word := "ABCCED"

		res := Exist(testBoard, word)
		So(res, ShouldBeTrue)
	})
}

func TestExist2(t *testing.T) {
	Convey("2", t, func() {
		testBoard := [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('C'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}
		word := "SEE"

		res := Exist(testBoard, word)
		So(res, ShouldBeTrue)
	})
}

func TestExist3(t *testing.T) {
	Convey("3", t, func() {
		testBoard := [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('C'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}
		word := "ABCB"

		res := Exist(testBoard, word)
		So(res, ShouldBeFalse)
	})
}

func TestExist4(t *testing.T) {
	Convey("4", t, func() {
		testBoard := [][]byte{
			{byte('a'), byte('a'), byte('b'), byte('b')},
			{byte('a'), byte('a'), byte('b'), byte('b')},
		}
		word := "baa"

		res := Exist(testBoard, word)
		So(res, ShouldBeTrue)
	})
}

func TestExist5(t *testing.T) {
	Convey("5", t, func() {
		testBoard := [][]byte{
			{byte('b'), byte('a'), byte('b'), byte('b'), byte('a')},
			{byte('a'), byte('a'), byte('b'), byte('a'), byte('b')},
			{byte('a'), byte('a'), byte('a'), byte('b'), byte('b')},
		}
		word := "aa"

		res := Exist(testBoard, word)
		So(res, ShouldBeTrue)
	})
}

func TestExist6(t *testing.T) {
	Convey("6", t, func() {
		testBoard := [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('E'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}
		word := "ABCESEEEFS"

		res := Exist(testBoard, word)
		So(res, ShouldBeTrue)
	})
}

package dfs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterCombinations(t *testing.T) {
	Convey("1", t, func() {
		res := LetterCombinations("234")
		exp := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}
		So(res, ShouldResemble, exp)
	})
}

func TestPermute(t *testing.T) {
	Convey("1", t, func() {
		input := []int{1, 2, 3}
		res := Permute(input)
		exp := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		So(res, ShouldResemble, exp)
	})
}

func TestPermute1(t *testing.T) {
	Convey("1", t, func() {
		input := []int{5, 4, 6, 2}
		res := Permute(input)
		exp := [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}}
		So(res, ShouldResemble, exp)
	})
}

func TestPermute2(t *testing.T) {
	Convey("1", t, func() {
		input := []int{5, 4, 6, 2}
		res := Permute2(input)
		exp := [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}}
		So(res, ShouldResemble, exp)
	})
}

func TestPermute3(t *testing.T) {
	Convey("1", t, func() {
		input := []int{1, 1, 2}
		res := PermuteUnique(input)
		exp := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
		So(res, ShouldResemble, exp)
	})
}

func TestSubsets(t *testing.T) {
	Convey("1", t, func() {
		input := []int{1, 2, 3}
		res := subsets(input)
		exp := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}
		So(res, ShouldEqual, exp)
	})
}

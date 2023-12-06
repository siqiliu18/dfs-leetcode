package isbipartite

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsBipartite1(t *testing.T) {
	Convey("1", t, func() {
		graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
		res := IsBipartite(graph)
		So(res, ShouldBeTrue)
	})
}

func TestIsBipartite2(t *testing.T) {
	Convey("1", t, func() {
		graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
		res := IsBipartite(graph)
		So(res, ShouldBeTrue)
	})
}

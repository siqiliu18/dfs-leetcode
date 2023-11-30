package accountsmerge

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountsMerge(t *testing.T) {
	Convey("1", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"}}
		res := AccountsMergeNotWorking(accounts)
		So(res, ShouldEqual, [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}})
	})
	Convey("2", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "john00@mail.com", "johnnybravo@mail.com"},
		}
		res := AccountsMergeNotWorking(accounts)
		So(res, ShouldEqual, [][]string{})
	})
}

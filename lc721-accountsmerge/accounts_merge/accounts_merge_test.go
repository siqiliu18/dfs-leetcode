package accountsmerge

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountsMergeNotWorking(t *testing.T) {
	Convey("1", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"}}
		res := AccountsMergeNotWorking(accounts)
		So(res, ShouldEqual, [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
		})
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
		So(res, ShouldEqual, [][]string{
			{"Mary", "mary@mail.com"},
			{"John", "john00@mail.com", "john_newyork@mail.com", "johnnybravo@mail.com", "johnsmith@mail.com"},
		})
	})
}

func TestAccountsMergeBFS(t *testing.T) {
	Convey("1", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"}}
		res := AccountsMergeBFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
		})
	})
	Convey("2", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "johnnybravo@mail.com", "johnnybravo1@mail.com"},
		}
		res := AccountsMergeBFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com", "johnnybravo1@mail.com"},
		})
	})
	Convey("3", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "john00@mail.com", "johnnybravo@mail.com"},
			{"John", "john00@mail.com", "john01@mail.com"},
		}
		res := AccountsMergeBFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"Mary", "mary@mail.com"},
			{"John", "john00@mail.com", "john_newyork@mail.com", "johnnybravo@mail.com", "johnsmith@mail.com"},
		})
	})
}

func TestAccountsMergeBFS2(t *testing.T) {
	Convey("1", t, func() {
		accounts := [][]string{
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
			{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
			{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
			{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
			{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
		}
		res := AccountsMergeBFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
			{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
			{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
			{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
		})
	})
}

func TestAccountsMergeDFS1(t *testing.T) {
	Convey("1", t, func() {
		accounts := [][]string{
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
			{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
			{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
			{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
			{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
		}
		res := AccountsMergeDFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
			{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
			{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
			{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
		})
	})
}

func TestAccountsMergeDFS2(t *testing.T) {
	Convey("1", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"}}
		res := AccountsMergeDFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
		})
	})
	Convey("2", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "johnnybravo@mail.com", "johnnybravo1@mail.com"},
		}
		res := AccountsMergeDFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com", "johnnybravo1@mail.com"},
		})
	})
	Convey("3", t, func() {
		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "john00@mail.com", "johnnybravo@mail.com"},
			{"John", "john00@mail.com", "john01@mail.com"},
		}
		res := AccountsMergeDFS(accounts)
		So(res, ShouldResemble, [][]string{
			{"Mary", "mary@mail.com"},
			{"John", "john00@mail.com", "john_newyork@mail.com", "johnnybravo@mail.com", "johnsmith@mail.com"},
		})
	})
}

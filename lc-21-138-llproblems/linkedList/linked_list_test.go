package linkedList

import (
	"dfs_leetcode/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoLists(t *testing.T) {
	Convey("1", t, func() {
		// list1 := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 4, Next: nil}}}
		// list2 := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: nil}}}
		list1 := utils.BuildLL("[1,2,4]")
		list2 := utils.BuildLL("[1,3,4]")
		res := MergeTwoLists(list1, list2)
		resArr := utils.ConvertLlToArr(res)
		exp := utils.BuildLL("1,1,2,3,4,5")
		expArr := utils.ConvertLlToArr(exp)
		So(resArr, ShouldResemble, expArr)
	})
}

func TestReorderList1(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4]")
		ReorderList(head)
		So(head.Next.Val, ShouldEqual, 4)
		So(head.Next.Next.Val, ShouldEqual, 2)
		So(head.Next.Next.Next.Val, ShouldEqual, 3)
	})
}

func TestReorderList2(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4,5]")
		ReorderList(head)
		So(head.Next.Val, ShouldEqual, 5)
		So(head.Next.Next.Val, ShouldEqual, 2)
		So(head.Next.Next.Next.Val, ShouldEqual, 4)
	})
}

func TestReorderList3(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4,5,6]")
		ReorderList(head)
		So(head.Next.Val, ShouldEqual, 6)
		So(head.Next.Next.Val, ShouldEqual, 2)
		So(head.Next.Next.Next.Val, ShouldEqual, 5)
	})
}

func TestReorderList4(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2]")
		ReorderList(head)
		So(head.Val, ShouldEqual, 1)
		So(head.Next.Val, ShouldEqual, 2)
	})
}

func TestReorderList5(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4,5,6,7]")
		ReorderList(head)
		So(head.Next.Val, ShouldEqual, 7)
		So(head.Next.Next.Val, ShouldEqual, 2)
		So(head.Next.Next.Next.Val, ShouldEqual, 6)
		So(head.Next.Next.Next.Next.Next.Val, ShouldEqual, 5)
	})
}

func TestReorderList6(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4,5,6,7,8,9,10,11]")
		ReorderList(head)
		arr := utils.ConvertLlToArr(head)
		lastIndex := len(arr) - 1
		So(arr, ShouldResemble, []int{1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6})
		So(arr[lastIndex-1], ShouldEqual, 7)
		So(arr[lastIndex], ShouldEqual, 6)
	})
}

func TestReorderList7(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4,5,6,7,8,9,10,11,12]")
		ReorderList(head)
		arr := utils.ConvertLlToArr(head)
		So(arr, ShouldResemble, []int{1, 12, 2, 11, 3, 10, 4, 9, 5, 8, 6, 7})
	})
}

func TestReorderList8(t *testing.T) {
	Convey("1", t, func() {
		head := utils.BuildLL("[1,2,3,4,5,6,7,8,9,10,11,12,13,14]")
		ReorderList(head)
		arr := utils.ConvertLlToArr(head)
		So(arr, ShouldResemble, []int{1, 14, 2, 13, 3, 12, 4, 11, 5, 10, 6, 9, 7, 8})
	})
}

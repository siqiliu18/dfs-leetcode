package utils

import (
	"log"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func BuildLL(llStr string) *ListNode {
	// [1,2,4]
	numsComas := strings.Trim(llStr, "[]")
	numsStrArr := strings.Split(numsComas, ",")

	head := &ListNode{Val: 0, Next: nil}
	mov := head

	for _, numStr := range numsStrArr {
		num, err := strconv.ParseInt(numStr, 10, 0)
		if err != nil {
			log.Fatalf("%v is not a number", num)
		}
		mov.Next = &ListNode{int(num), nil}
		mov = mov.Next
	}
	return head.Next
}

func ConvertLlToArr(ll *ListNode) []int {
	res := []int{}
	for ; ll != nil; ll = ll.Next {
		res = append(res, ll.Val)
	}
	return res
}

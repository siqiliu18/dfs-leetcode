package linkedList

import (
	"dfs_leetcode/utils"
)

// lc 21
func MergeTwoLists(list1, list2 *utils.ListNode) *utils.ListNode {
	head := &utils.ListNode{Val: 0, Next: nil}
	it := head

	for ; list1 != nil || list2 != nil; it = it.Next {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				it.Next = &utils.ListNode{Val: list1.Val, Next: nil}
				list1 = list1.Next
			} else {
				it.Next = &utils.ListNode{Val: list2.Val, Next: nil}
				list2 = list2.Next
			}
		} else if list1 != nil && list2 == nil {
			it.Next = &utils.ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else if list1 == nil && list2 != nil {
			it.Next = &utils.ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}
	}

	return head.Next
}

// lc 138
// use a map<Node, Node> to record old node to new node under the same val,
// when iterate through new list, create 2 its
// new it random = map[old it ramdom]
// ex: map[old 13 random] = 7
func CopyRandomList(head *utils.Node) *utils.Node {
	if head == nil {
		return nil
	}

	it := head

	nodeMap := make(map[*utils.Node]*utils.Node)

	for it != nil {
		nodeMap[it] = &utils.Node{Val: it.Val}
		it = it.Next
	}
	it = head
	for it != nil {
		nodeMap[it].Random = nodeMap[it.Random]
		nodeMap[it].Next = nodeMap[it.Next]
		it = it.Next
	}
	return nodeMap[head]
}

// lc 143
func ReorderList(head *utils.ListNode) {
	// create a map for reverse pointers
	nodeMap := make(map[int]*utils.ListNode)

	n := 0
	mov := head
	for mov != nil {
		nodeMap[n] = mov
		n += 1
		mov = mov.Next
	}
	if len(nodeMap) <= 2 {
		return
	}

	mapLen := len(nodeMap)
	n = mapLen - 1 // now n is the last index of the ll
	mov = head

	for mov.Next != nil && mov.Next.Next != nil {
		nodeMap[n].Next = mov.Next // point last node to the next node of current
		mov.Next = nodeMap[n]      // point current node to new node: 1 -> 5 -> 2
		nodeMap[n-1].Next = nil
		mov = mov.Next.Next
		n -= 1
	}
}

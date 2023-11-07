package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var nullStr string = "null"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeFromStringArray(strArr string) *TreeNode {
	withoutBlackets := strings.Trim(strArr, "[]")
	arrOfStr := strings.Split(withoutBlackets, ",")
	fmt.Println(arrOfStr)

	root := new(TreeNode)

	if len(arrOfStr) > 0 && arrOfStr[0] != nullStr {
		num, err := strconv.ParseInt(arrOfStr[0], 10, 0)
		if err != nil {
			log.Fatalf("the first element %v is not a number", arrOfStr[0])
		}
		root.Val = int(num)

		dfs(arrOfStr, 0, 1, root)
	}

	printPreOrder(root)
	fmt.Println()
	printPostOrder(root)
	fmt.Println()
	printInOrder(root)

	return root
}

func dfs(arrOfStr []string, index int, diff int, root *TreeNode) {
	if root == nil {
		return
	}

	leftChildIndex := index + diff
	rightChildIndex := index + diff + 1

	if leftChildIndex < len(arrOfStr) && arrOfStr[leftChildIndex] != nullStr {
		num, err := strconv.ParseInt(arrOfStr[leftChildIndex], 10, 0)
		if err != nil {
			log.Fatalf("the left child in arrary %v is not a number", arrOfStr[leftChildIndex])
		}
		root.Left = &TreeNode{
			Val: int(num),
		}
		dfs(arrOfStr, leftChildIndex, leftChildIndex+1, root.Left)
	}

	if rightChildIndex < len(arrOfStr) && arrOfStr[rightChildIndex] != nullStr {
		num, err := strconv.ParseInt(arrOfStr[rightChildIndex], 10, 0)
		if err != nil {
			log.Fatalf("the right child in array %v is not a number", arrOfStr[rightChildIndex])
		}
		root.Right = &TreeNode{
			Val: int(num),
		}
		dfs(arrOfStr, rightChildIndex, rightChildIndex+1, root.Right)
	}
}

func printPreOrder(root *TreeNode) {
	if root == nil {
		return
	} else {
		fmt.Printf("%d ", root.Val)
		printPreOrder(root.Left)
		printPreOrder(root.Right)
	}
}

func printPostOrder(root *TreeNode) {
	if root == nil {
		return
	} else {
		printPostOrder(root.Left)
		printPostOrder(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	} else {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		printInOrder(root.Right)
	}
}

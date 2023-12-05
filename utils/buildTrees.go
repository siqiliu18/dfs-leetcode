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

func BuildTreeFromStringArrayForLoop(strArr string) *TreeNode {
	withoutBlackets := strings.Trim(strArr, "[]")
	arrOfStr := strings.Split(withoutBlackets, ",")
	arrLen := len(arrOfStr)

	indexMap := make(map[int][]int)

	diff := 0
	for currIndex, str := range arrOfStr {
		if str != nullStr {
			diff++
			leftChildIndex := currIndex + diff
			rightChildIndex := currIndex + diff + 1
			if leftChildIndex <= arrLen && rightChildIndex <= arrLen {
				indexMap[currIndex] = append(indexMap[currIndex], currIndex+diff)
				indexMap[currIndex] = append(indexMap[currIndex], currIndex+diff+1)
			}
		} else {
			diff--
		}
	}

	root := new(TreeNode)

	if len(arrOfStr) > 0 && arrOfStr[0] != nullStr {
		num, err := strconv.ParseInt(arrOfStr[0], 10, 0)
		if err != nil {
			log.Fatalf("the first element %v is not a number", arrOfStr[0])
		}
		root.Val = int(num)
		dfsForLoop(root, 0, arrOfStr, indexMap)
	}

	fmt.Printf("PreOrder: ")
	printPreOrder(root)
	fmt.Println()
	fmt.Printf("PostOrder: ")
	printPostOrder(root)
	fmt.Println()
	fmt.Printf("InOrder: ")
	printInOrder(root)
	fmt.Println()

	return root
}

func dfsForLoop(root *TreeNode, index int, arrOfStr []string, indexMap map[int][]int) {
	if index >= len(arrOfStr) || root == nil {
		return
	}

	if _, ok := indexMap[index]; ok && len(indexMap[index]) > 0 {

		if indexMap[index][0] < len(arrOfStr) && arrOfStr[indexMap[index][0]] != nullStr {
			leftVal, _ := strconv.ParseInt(arrOfStr[indexMap[index][0]], 10, 0)
			root.Left = &TreeNode{
				Val: int(leftVal),
			}
		}

		if indexMap[index][1] < len(arrOfStr) && arrOfStr[indexMap[index][1]] != nullStr {
			rightVal, _ := strconv.ParseInt(arrOfStr[indexMap[index][1]], 10, 0)
			root.Right = &TreeNode{
				Val: int(rightVal),
			}
		}

		dfsForLoop(root.Left, indexMap[index][0], arrOfStr, indexMap)
		dfsForLoop(root.Right, indexMap[index][1], arrOfStr, indexMap)
	}
}

/*
	 this method has a bug for this case [5,4,8,11,null,13,4,7,2,null,null,null,1]
		which it expects all positions in all level need to be fill out with null string for nil nodes
		but leetcode don't add extra nulls under a null node
*/
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

	fmt.Printf("PreOrder: ")
	printPreOrder(root)
	fmt.Println()
	fmt.Printf("PostOrder: ")
	printPostOrder(root)
	fmt.Println()
	fmt.Printf("InOrder: ")
	printInOrder(root)
	fmt.Println()

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

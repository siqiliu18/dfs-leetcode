package main

import (
	"fmt"

	existfun "dfs_leetcode79/dfs"
)

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Input: board = [
	["A","B","C","E"],
	["S","F","C","S"],
	["A","D","E","E"]
	],
	word = "ABCCED"
Output: true

Input: board = [
	["A","B","C","E"],
	["S","F","C","S"],
	["A","D","E","E"]
	],
	word = "SEE"
Output: true

Input: board = [
	["A","B","C","E"],
	["S","F","C","S"],
	["A","D","E","E"]
	],
	word = "ABCB"
Output: false
*/

func main() {
	fmt.Println("Hello World!")

	printBoard := [][]byte{
		{byte('A'), byte('B'), byte('C'), byte('E')},
		{byte('S'), byte('F'), byte('C'), byte('S')},
		{byte('A'), byte('D'), byte('E'), byte('E')},
	}
	res := existfun.Exist(printBoard, "")
	fmt.Println(res)
}

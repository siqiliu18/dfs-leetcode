package main

import (
	"dfs_leetcode/lc200-numofislands/dfs"
	"fmt"
)

func main() {

	grid1 := [][]byte{
		{byte('1'), byte('1'), byte('1'), byte('1'), byte('0')},
		{byte('1'), byte('1'), byte('0'), byte('1'), byte('0')},
		{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
		{byte('0'), byte('0'), byte('0'), byte('0'), byte('0')},
	}

	fmt.Println(dfs.NumIslands(grid1))

	grid2 := [][]byte{
		{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
		{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
		{byte('0'), byte('0'), byte('1'), byte('0'), byte('0')},
		{byte('0'), byte('0'), byte('0'), byte('1'), byte('1')},
	}

	fmt.Println(dfs.NumIslands(grid2))
}

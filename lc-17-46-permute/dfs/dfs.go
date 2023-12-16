package dfs

/*
	Leetcode problem 17:

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]
*/
var dMap map[byte]string = map[byte]string{
	byte('1'): "",
	byte('2'): "abc",
	byte('3'): "def",
	byte('4'): "ghi",
	byte('5'): "jkl",
	byte('6'): "mno",
	byte('7'): "pqrs",
	byte('8'): "tuv",
	byte('9'): "wxyz",
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := new([]string)
	combStr := make([]byte, len(digits))
	combRec(&digits, 0, combStr, res)
	return *res
}

func combRec(digits *string, index int, combStr []byte, outVec *[]string) {
	if index == len(*digits) {
		*outVec = append(*outVec, string(combStr))
		return
	}

	tmpStr := dMap[(*digits)[index]]
	for i := 0; i < len(tmpStr); i++ {
		combStr[index] = tmpStr[i]
		combRec(digits, index+1, combStr, outVec)
	}
}

/*
46. Permutations

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Explanation 1: https://www.youtube.com/watch?v=H232aocj7bQ (I think this one is better)
Explanation 2: https://www.youtube.com/watch?v=Nabbpl7y4Lo&t=506s
*/
func Permute(nums []int) [][]int {
	res := new([][]int)
	combArr := []int{}
	visited := make([]bool, len(nums))
	permRec(nums, combArr, visited, res)
	return *res
}

func permRec(nums []int, combArr []int, visited []bool, res *[][]int) {
	if len(combArr) == len(nums) {
		// *res = append(*res, combArr)
		/*
			it is interesting that by direct append *res = append(*res, combArr)
			the res will be modified!!! See {5,4,6,2} example.
			{{5,4,2,6},{5,4,2,6},...}
		*/
		// cp := make([]int, len(combArr))
		// copy(cp, combArr) // deep copy
		*res = append(*res, append([]int{}, combArr...)) // create new slice
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		combArr = append(combArr, nums[i])
		permRec(nums, combArr, visited, res)
		combArr = combArr[:len(combArr)-1]
		visited[i] = false
	}
}

func permRec2(nums []int, combArr []int, res *[][]int) {
	if len(combArr) == len(nums) {
		*res = append(*res, append([]int{}, combArr...))
		return
	}
	for _, num := range nums {
		if contains(combArr, num) {
			continue
		}
		combArr = append(combArr, num)
		permRec2(nums, combArr, res)
		combArr = combArr[:len(combArr)-1]
	}
}

func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}

	return false
}

// why this one no need deep copy??????
func Permute2(nums []int) [][]int {
	o := [][]int{}
	curHash := make(map[int]bool)
	backtracking(nums, &o, []int{}, curHash)
	return o
}

func backtracking(nums []int, o *[][]int, cur []int, curHash map[int]bool) {
	if len(nums) == len(cur) {
		*o = append(*o, cur)
		return
	}
	// for each value in nums, start with nums[i]
	for i := 0; i < len(nums); i++ {
		// record existing values in hash, and skip if already used
		if curHash[nums[i]] == true {
			continue
		}
		curHash[nums[i]] = true
		// add the next nums[i] value to current list
		cur := append(cur, nums[i]) // !!!!!! := !!!!!!!!!!!!
		// recursively call function
		backtracking(nums, o, cur, curHash)
		// remove last item from cur
		delete(curHash, cur[len(cur)-1])
		cur = cur[:len(cur)-1]
	}
}

/* ex
package main

import (
	"fmt"
)

func main() {
	cur := []int{1,2}
	fmt.Println(cur)
	for i := 3; i < 6; i++ {
    	cur := append(cur, i) //// !!!! it is OK to do it inside a for loop block,
    	fmt.Println(cur)
  	}
}

output:
[1 2]
[1 2 3]
[1 2 4]
[1 2 5]
*/

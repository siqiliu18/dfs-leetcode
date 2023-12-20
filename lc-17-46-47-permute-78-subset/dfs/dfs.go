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

func permute(nums []int) [][]int {
	res := new([][]int)
	combArr := []int{}
	visited := make([]bool, len(nums))
	permRec(nums, combArr, visited, res)
	return *res
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

/*
	Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

Input: nums = [1,1,2]
Output:
[[1,1,2],

	[1,2,1],
	[2,1,1]]

C++ code, use a set to check dups
class Solution {
public:

	    vector<vector<int>> permuteUnique(vector<int>& nums) {
	        vector<vector<int>> res;
	        vector<int> pvec;
	        vector<bool> visited (nums.size(), false);
	        // backtrackVisited(res, pvec, nums, visited);
	        // backtrackSwap(res, nums, 0);
	        backtrackVisitedWithDuplicatedNums(res, pvec, nums, visited);
	        return res;
	    }

	    void backtrackVisitedWithDuplicatedNums(vector<vector<int>>& res, vector<int>& pvec, vector<int>& nums, vector<bool>& visited) {
	        // base case
	        if (pvec.size() == nums.size()) {
	            res.push_back(pvec);
	            return;
	        }

	        set<int> checkDup;
	        for (int i = 0; i < nums.size(); ++i) {
	            if (visited[i] || checkDup.count(nums[i])) continue;
	            checkDup.insert(nums[i]);
	            visited[i] = true;
	            pvec.push_back(nums[i]);
	            backtrackVisitedWithDuplicatedNums(res, pvec, nums, visited);
	            pvec.pop_back();
	            visited[i] = false;
	        }
	    }
	};
*/
func PermuteUnique(nums []int) [][]int {
	res := new([][]int)
	combArr := []int{}
	visited := make([]bool, len(nums))
	PermuteUniqueRecur(nums, combArr, res, visited)
	return *res
}

func PermuteUniqueRecur(nums []int, combArr []int, res *[][]int, visited []bool) {
	if len(combArr) == len(nums) {
		*res = append(*res, append([]int{}, combArr...))
		return
	}
	checkDup := make(map[int]int)
	for i, num := range nums {
		if _, ok := checkDup[num]; ok || visited[i] {
			continue
		}
		checkDup[num] = 0
		visited[i] = true
		combArr = append(combArr, num)
		PermuteUniqueRecur(nums, combArr, res, visited)
		combArr = combArr[:len(combArr)-1]
		visited[i] = false
	}
}

/*
	78.Subsets

Given an integer array nums of unique elements, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.


Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]
*/

// method 1: create a count that tells how many nums will be inserted.
//   - at the count loop, start is always from 0
//   - at the recurrsion loop, start is 1 index to the right of current index
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for count := 1; count <= len(nums); count++ {
		curArr := []int{}
		subsetsRecur(nums, count, curArr, &res, 0)
	}
	return res
}

func subsetsRecur(nums []int, count int, curArr []int, res *[][]int, start int) {
	if len(curArr) == count {
		subsetCopy := make([]int, len(curArr))
		copy(subsetCopy, curArr)
		*res = append(*res, curArr)
		return
	}
	for i := start; i < len(nums); i++ {
		curArr = append(curArr, nums[i])
		subsetsRecur(nums, count, curArr, res, i+1)
		curArr = curArr[:len(curArr)-1]
	}
}

// UPDATE!!!: https://www.youtube.com/watch?v=3tpjp5h3M6Y
func subsets2(nums []int) [][]int {
	res := [][]int{{}}
	// for count := 1; count <= len(nums); count++ {
	curArr := []int{}
	subsetsRecur2(nums, curArr, &res, 0)
	// }
	return res
}

func subsetsRecur2(nums []int, current []int, res *[][]int, start int) {
	subsetCopy := make([]int, len(current))
	copy(subsetCopy, current)
	*res = append(*res, subsetCopy)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		subsetsRecur2(nums, current, res, i+1)
		current = current[:len(current)-1]
	}
}

// method 2: create res = [[]]
//   - loop through res, make a copy of each element and add current num into each
//   - nums = [1,2,3]
//   - make a copy : [[], []], then add 1 into the copy => [[], [1]]
//   - make a copy [[], [1], [], [1]], add 2 into copies => [[], [1], [2], [1, 2]]
func subsets3(nums []int) [][]int {
	res := [][]int{{}}
	resLen := len(res)
	for _, num := range nums {
		for i := 0; i < resLen; i++ {
			curArr := res[i]
			curArr = append(curArr, num)
			res = append(res, curArr)
		}
		resLen = len(res)
	}
	return res
}

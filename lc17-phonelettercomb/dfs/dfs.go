package dfs

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

func Permute(nums []int) [][]int {
	res := new([][]int)
	combArr := []int{}
	visited := make([]bool, len(nums))
	permRec(nums, combArr, visited, res)
	return *res
}

func permRec(nums []int, combArr []int, visited []bool, res *[][]int) {
	if len(combArr) == len(nums) {
		*res = append(*res, combArr)
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
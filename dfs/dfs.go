package dfs

type Node struct {
	x int
	y int
}

func DFS(board [][]byte, node Node, word string, index int, visited map[Node]int, rowLen int, colLen int) bool {
	// if current node == letter, extend to max 4 directions
	// if current node != letter, return false

	if index == len(word)-1 {
		return true
	}

	d1, d2, d3, d4 := false, false, false, false

	// left node
	leftNode := Node{
		x: node.x,
		y: node.y - 1,
	}

	if _, ok := visited[leftNode]; !ok && node.y-1 >= 0 && board[node.x][node.y-1] == word[index+1] {
		visited[leftNode] = 0
		d1 = DFS(board, leftNode, word, index+1, visited, rowLen, colLen)
		if !d1 {
			delete(visited, leftNode)
		}
	}

	// right node
	rightNode := Node{
		x: node.x,
		y: node.y + 1,
	}
	if _, ok := visited[rightNode]; !ok && node.y+1 < colLen && board[node.x][node.y+1] == word[index+1] {
		visited[rightNode] = 0
		d2 = DFS(board, rightNode, word, index+1, visited, rowLen, colLen)
		if !d2 {
			delete(visited, rightNode)
		}
	}

	// top node
	topNode := Node{
		x: node.x - 1,
		y: node.y,
	}
	if _, ok := visited[topNode]; !ok && node.x-1 >= 0 && board[node.x-1][node.y] == word[index+1] {
		visited[topNode] = 0
		d3 = DFS(board, topNode, word, index+1, visited, rowLen, colLen)
		if !d3 {
			delete(visited, topNode)
		}
	}

	// bottom node
	bottomNode := Node{
		x: node.x + 1,
		y: node.y,
	}
	if _, ok := visited[bottomNode]; !ok && node.x+1 < rowLen && board[node.x+1][node.y] == word[index+1] {
		visited[bottomNode] = 0
		d4 = DFS(board, bottomNode, word, index+1, visited, rowLen, colLen)
		if !d4 {
			delete(visited, bottomNode)
		}
	}

	return d1 || d2 || d3 || d4
}

func Exist(board [][]byte, word string) bool {
	rowLen := len(board)
	colLen := len(board[0])
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if board[row][col] == word[0] { // back for word len = 0 case
				firstNode := Node{
					x: row,
					y: col,
				}
				visited := make(map[Node]int)
				visited[firstNode] = 0
				if DFS(board, firstNode, word, 0, visited, rowLen, colLen) {
					return true
				}
			}
		}
	}
	return false
}

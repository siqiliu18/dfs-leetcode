package utils

import (
	"log"
	"strconv"
	"strings"
)

// This is related to leetcode problem 133 https://leetcode.com/problems/clone-graph/submissions/547251836/

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func BuildGraphsBFS(tableStr string) *GraphNode {
	// [[2,4],[1,3],[2,4],[1,3]]
	tableRows := strings.Split(tableStr, "],[")
	treeTableRowLen := len(tableRows)
	treeTable := make([][]int, treeTableRowLen)
	for row, str := range tableRows {
		withoutBlackets := strings.Trim(str, "[]")
		rowArr := strings.Split(withoutBlackets, ",")
		for _, numStr := range rowArr {
			num, err := strconv.ParseInt(numStr, 10, 0)
			if err != nil {
				log.Fatalf("the first element %v is not a number", numStr)
			}
			treeTable[row] = append(treeTable[row], int(num))
		}
	}

	indexMap := make(map[int][]int)
	for i := 0; i < len(treeTable); i++ {
		indexMap[i+1] = treeTable[i]
	}

	nodeMap := make(map[int]*GraphNode)

	startNode := &GraphNode{
		Val: 1,
	}

	nodeMap[1] = startNode

	nodeq := [](*GraphNode){}

	nodeq = append(nodeq, startNode)

	for len(nodeq) != 0 {
		currNode := nodeq[0]
		nodeq = nodeq[1:]

		for _, neighborNum := range indexMap[currNode.Val] {
			if _, ok := nodeMap[neighborNum]; !ok {
				newNode := &GraphNode{
					Val: neighborNum,
				}
				nodeMap[neighborNum] = newNode
				nodeq = append(nodeq, newNode)
			}
			nodeMap[currNode.Val].Neighbors = append(nodeMap[currNode.Val].Neighbors, nodeMap[neighborNum])
		}
	}

	return startNode
}

func BuildGraphsDFS(tableStr string) *GraphNode {
	// [[2,4],[1,3],[2,4],[1,3]]
	tableRows := strings.Split(tableStr, "],[")
	treeTableRowLen := len(tableRows)
	treeTable := make([][]int, treeTableRowLen)
	for row, str := range tableRows {
		withoutBlackets := strings.Trim(str, "[]")
		rowArr := strings.Split(withoutBlackets, ",")
		for _, numStr := range rowArr {
			num, err := strconv.ParseInt(numStr, 10, 0)
			if err != nil {
				log.Fatalf("the first element %v is not a number", numStr)
			}
			treeTable[row] = append(treeTable[row], int(num))
		}
	}

	indexMap := make(map[int][]int)
	for i := 0; i < len(treeTable); i++ {
		indexMap[i+1] = treeTable[i]
	}

	nodeMap := make(map[int]*GraphNode)

	startNode := &GraphNode{
		Val: 1,
	}

	nodeMap[1] = startNode

	toGo := make(map[int]bool)

	dfsGraphs(1, indexMap, nodeMap, toGo)

	return startNode
}

func dfsGraphs(currVal int, indexMap map[int][]int, nodeMap map[int]*GraphNode, toGo map[int]bool) {
	toGo[currVal] = false
	for _, neighborNum := range indexMap[currVal] {
		if _, ok := nodeMap[neighborNum]; !ok {
			newNode := &GraphNode{
				Val: neighborNum,
			}
			nodeMap[neighborNum] = newNode
			toGo[neighborNum] = true
		}
		nodeMap[currVal].Neighbors = append(nodeMap[currVal].Neighbors, nodeMap[neighborNum])
	}

	for val, togo := range toGo {
		if togo {
			dfsGraphs(val, indexMap, nodeMap, toGo)
		}
	}
}

/* This is more clear
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;
    Node() {
        val = 0;
        neighbors = vector<Node*>();
    }
    Node(int _val) {
        val = _val;
        neighbors = vector<Node*>();
    }
    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

class Solution {
	public:
		Node* cloneGraph(Node* node) {
			if (!node) {
				return NULL;
			}
			if (copies.find(node) == copies.end()) {
				copies[node] = new Node(node -> val, {});
				for (Node* neighbor : node -> neighbors) {
					copies[node] -> neighbors.push_back(cloneGraph(neighbor));
				}
			}
			return copies[node];
		}
	private:
		unordered_map<Node*, Node*> copies;
	};
*/

var NodeMap map[int]*GraphNode

func BuildGraphBackTrack(tableStr string) *GraphNode {
	// [[2,4],[1,3],[2,4],[1,3]]
	tableRows := strings.Split(tableStr, "],[")
	treeTableRowLen := len(tableRows)
	treeTable := make([][]int, treeTableRowLen)
	for row, str := range tableRows {
		withoutBlackets := strings.Trim(str, "[]")
		rowArr := strings.Split(withoutBlackets, ",")
		for _, numStr := range rowArr {
			num, err := strconv.ParseInt(numStr, 10, 0)
			if err != nil {
				log.Fatalf("the first element %v is not a number", numStr)
			}
			treeTable[row] = append(treeTable[row], int(num))
		}
	}

	indexMap := make(map[int][]int)
	for i := 0; i < len(treeTable); i++ {
		indexMap[i+1] = treeTable[i]
	}

	NodeMap = make(map[int]*GraphNode)

	return dfsBackTrack(1, indexMap)
}

func dfsBackTrack(currVal int, indexMap map[int][]int) *GraphNode {
	if _, ok := NodeMap[currVal]; !ok {
		NodeMap[currVal] = &GraphNode{
			Val: currVal,
		}
		for _, neighborNum := range indexMap[currVal] {
			NodeMap[currVal].Neighbors = append(NodeMap[currVal].Neighbors, dfsBackTrack(neighborNum, indexMap))
		}
	}
	return NodeMap[currVal]
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("clone graph")
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodeMap := make(map[*Node]*Node)

	newRoot := &Node{
		Val: node.Val,
	}

	nodeMap[node] = newRoot

	nodeq := [](*Node){}
	nodeq = append(nodeq, node)

	for len(nodeq) != 0 {
		currNode := nodeq[0]
		nodeq = nodeq[1:]

		for _, neighbor := range currNode.Neighbors {
			if _, ok := nodeMap[neighbor]; !ok {
				newNode := &Node{
					Val: neighbor.Val,
				}
				nodeMap[neighbor] = newNode
				nodeq = append(nodeq, neighbor)
			}
			nodeMap[currNode].Neighbors = append(nodeMap[currNode].Neighbors, nodeMap[neighbor])
		}
	}

	return newRoot
}

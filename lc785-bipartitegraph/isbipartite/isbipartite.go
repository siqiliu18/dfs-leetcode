package isbipartite

func IsBipartite(graph [][]int) bool {
	n := len(graph) // number of nodes in graph
	visited := make([]bool, n)
	s1 := make([]bool, n)
	s2 := make([]bool, n)

	for i := range graph {
		if visited[i] {
			continue
		}

		s1[i] = true
		queue := []int{i}

		for len(queue) != 0 {
			u := queue[0]
			queue = queue[1:]
			visited[u] = true
			for _, v := range graph[u] {
				if !visited[v] {
					queue = append(queue, v)
					if s1[u] {
						s2[v] = true
					} else {
						s1[v] = true
					}
				}
			}
		}
	}

	for i := range s1 {
		if s1[i] == s2[i] {
			return false
		}
	}
	return len(s1) == len(s2)
}

// https://www.youtube.com/watch?v=670Gn4e89B8
func IsBipartite2(graph [][]int) bool {
	n := len(graph) // number of nodes in graph
	visited := make([]int, n)

	for i := range graph {
		if visited[i] != 0 {
			continue
		}
		queue := []int{i}
		visited[i] = 1
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			curLable := visited[cur]
			neighborLable := setNeighborLable(curLable)
			for _, neighbor := range graph[cur] {
				if visited[neighbor] == 0 {
					visited[neighbor] = neighborLable
					queue = append(queue, neighbor)
				} else if visited[neighbor] != neighborLable {
					return false
				}
			}
		}
	}
	return true
}

func setNeighborLable(curLable int) int {
	if curLable == 1 {
		return 2
	}
	return 1
}

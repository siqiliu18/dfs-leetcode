package dfs

import "math"

type Coordi struct {
	x   int
	y   int
	val int
}

func ShortestDistance(grid [][]int) int {
	height := len(grid)
	wide := len(grid[0])

	grid2 := make([][]int, height)

	ones := []Coordi{}

	for x := 0; x < height; x++ {
		grid2[x] = make([]int, wide)
		for y := 0; y < wide; y++ {
			grid2[x][y] = 0
			if grid[x][y] == 1 {
				ones = append(ones, Coordi{x, y, 1})
			}
		}
	}

	for i, oneCoordi := range ones {
		bfs(oneCoordi, grid, grid2, height, wide, (i+1)*-1)
	}

	min := math.MaxInt
	for x := 0; x < height; x++ {
		for y := 0; y < wide; y++ {
			if grid[x][y] != 1 && grid[x][y] != 2 {
				if grid2[x][y] < min {
					min = grid2[x][y]
				}
			}
		}
	}

	if min == math.MaxInt || min == 0 {
		return -1
	}

	return min
}

func bfs(oneCoordi Coordi, grid, grid2 [][]int, height, wide, path int) {
	fourD := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Coordi{oneCoordi}

	for len(queue) != 0 {
		currCoor := queue[0]
		queue = queue[1:]

		for _, di := range fourD {
			if currCoor.x+di[0] >= 0 && currCoor.x+di[0] < height && currCoor.y+di[1] >= 0 && currCoor.y+di[1] < wide && grid[currCoor.x+di[0]][currCoor.y+di[1]] != 2 {
				if grid[currCoor.x+di[0]][currCoor.y+di[1]] != path && grid[currCoor.x+di[0]][currCoor.y+di[1]] != 1 && grid[currCoor.x+di[0]][currCoor.y+di[1]] != 2 {
					grid2[currCoor.x+di[0]][currCoor.y+di[1]] = currCoor.val + grid2[currCoor.x+di[0]][currCoor.y+di[1]]
					grid[currCoor.x+di[0]][currCoor.y+di[1]] = path
					queue = append(queue, Coordi{currCoor.x + di[0], currCoor.y + di[1], currCoor.val + 1})
				}
			}
		}
	}
}

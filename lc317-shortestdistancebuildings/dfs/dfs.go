package dfs

import (
	"math"
)

type Coordi struct {
	x   int
	y   int
	val int
}

func ShortestDistance(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	grid2 := make([][]int, rows)

	ones := []Coordi{}

	for x := 0; x < rows; x++ {
		grid2[x] = make([]int, cols)
		for y := 0; y < cols; y++ {
			grid2[x][y] = 0
			if grid[x][y] == 1 {
				ones = append(ones, Coordi{x, y, 1})
			}
		}
	}

	round := 0
	for i, oneCoordi := range ones {
		round = (i + 1) * -1
		move := bfs(oneCoordi, grid, grid2, rows, cols, round)
		if !move {
			return -1
		}
	}

	min := math.MaxInt
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] != 1 && grid[x][y] != 2 {
				if grid[x][y] != round {
					return -1
				}
				if grid2[x][y] < min {
					min = grid2[x][y]
				}
			}
		}
	}

	return min
}

func bfs(oneCoordi Coordi, grid, grid2 [][]int, rows, cols, round int) bool {
	fourD := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Coordi{oneCoordi}

	move := false // check if any move
	for len(queue) != 0 {
		currCoor := queue[0]
		queue = queue[1:]

		for _, di := range fourD {
			if currCoor.x+di[0] >= 0 && currCoor.x+di[0] < rows && currCoor.y+di[1] >= 0 && currCoor.y+di[1] < cols && grid[currCoor.x+di[0]][currCoor.y+di[1]] != 2 {
				if grid[currCoor.x+di[0]][currCoor.y+di[1]] != round && grid[currCoor.x+di[0]][currCoor.y+di[1]] != 1 && grid[currCoor.x+di[0]][currCoor.y+di[1]] != 2 {
					grid2[currCoor.x+di[0]][currCoor.y+di[1]] = currCoor.val + grid2[currCoor.x+di[0]][currCoor.y+di[1]]
					grid[currCoor.x+di[0]][currCoor.y+di[1]] = round
					queue = append(queue, Coordi{currCoor.x + di[0], currCoor.y + di[1], currCoor.val + 1})
					move = true
				}
			}
		}
	}
	return move
}

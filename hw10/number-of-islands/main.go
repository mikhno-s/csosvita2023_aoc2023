package main

import "fmt"

func main() {
	in1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(in1))
	in2 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(in2))
	in3 := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Println(numIslands(in3))
}

func numIslands(grid [][]byte) int {
	var islands int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' {
				dfs(grid, r, c)
				islands++
			}

		}
	}
	return islands
}

func dfs(grid [][]byte, r int, c int) {
	if r < 0 || r >= len(grid) || c >= len(grid[r]) || c < 0 || grid[r][c] == '0' {
		return
	}

	// mark as visited
	grid[r][c] = '0'

	// check right
	dfs(grid, r, c+1)
	// check left
	dfs(grid, r, c-1)
	// check bottom
	dfs(grid, r+1, c)
	// check top
	dfs(grid, r-1, c)

}

package main

import "fmt"

// type Tile struct {
// 	// Neighbours
// 	NG []*Tile
// }

func main() {

	// in1 := [][]byte{{1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}
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
	visited := map[*byte]bool{}
	var islands int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' && !visited[&grid[r][c]] {
				dfs(visited, grid, r, c)
				islands++
			}

		}
	}
	return islands
}

func dfs(visited map[*byte]bool, grid [][]byte, r int, c int) {
	if visited[&grid[r][c]] != true {
		visited[&grid[r][c]] = true
	} else {
		return
	}

	// check right
	if c+1 < len(grid[r]) && grid[r][c+1] == '1' && !visited[&grid[r][c+1]] {
		dfs(visited, grid, r, c+1)
	}
	// check left
	if c-1 >= 0 && grid[r][c-1] == '1' && !visited[&grid[r][c-1]] {
		dfs(visited, grid, r, c-1)
	}
	// check bottom
	if r+1 < len(grid) && grid[r+1][c] == '1' && !visited[&grid[r+1][c]] {
		dfs(visited, grid, r+1, c)
	}
	// check top
	if r-1 >= 0 && grid[r-1][c] == '1' && !visited[&grid[r-1][c]] {
		dfs(visited, grid, r-1, c)
	}
}

// Better solution that does not need additional memory
// func visit(grid [][]byte, i, j int) {
//     if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
//         return
//     }
//     grid[i][j] = '0'
//     visit(grid, i - 1, j)
//     visit(grid, i + 1, j)
//     visit(grid, i, j - 1)
//     visit(grid, i, j + 1)
// }

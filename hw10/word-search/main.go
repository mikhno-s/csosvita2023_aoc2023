package main

import "fmt"

// type Tile struct {
// 	// Neighbours
// 	NG []*Tile
// }

func main() {

	board1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board1, "ABCCED")) // true

	board2 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board2, "SEE")) // true

	board3 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board3, "ABCB")) // false

	board4 := [][]byte{
		{'a', 'a', 'b', 'a', 'a', 'b'},
		{'b', 'a', 'b', 'a', 'b', 'b'},
		{'b', 'a', 'b', 'b', 'b', 'b'},
		{'a', 'a', 'b', 'a', 'b', 'a'},
		{'b', 'b', 'a', 'a', 'a', 'b'},
		{'b', 'b', 'b', 'a', 'b', 'a'},
	}
	fmt.Println(exist(board4, "aaaababab")) // true

	board5 := [][]byte{
		{'a'},
	}
	fmt.Println(exist(board5, "a")) // false
}

func exist(board [][]byte, word string) bool {
	visited := map[*byte]bool{}
	for r := range board {
		for c := range board[r] {
			if board[r][c] == word[0] {
				if visit(visited, board, r, c, word, "") {
					return true
				}
			}
		}
	}
	return false
}

func visit(visited map[*byte]bool, board [][]byte, r int, c int, word string, current string) bool {
	if r >= len(board) || r < 0 || c >= len(board[r]) || c < 0 || visited[&board[r][c]] || len(current) == len(word) {
		return current == word
	}

	visited[&board[r][c]] = true
	current += string(board[r][c])

	good := true
	for i := range current {
		if current[i] != word[i] {
			good = false
			break
		}
	}
	if good {
		if visit(visited, board, r, c+1, word, current) ||
			visit(visited, board, r, c-1, word, current) ||
			visit(visited, board, r+1, c, word, current) ||
			visit(visited, board, r-1, c, word, current) {
			return true
		}
	}
	current = current[:len(current)-1]
	visited[&board[r][c]] = false

	return false
}

package main

import "fmt"

//四个方向
var dir  = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

//func exist(board [][]byte, word string) bool {
//	visited := make([][]bool, len(board))
//	//记录所有遍历过的元素
//	for i := 0; i < len(visited); i++ {
//		visited[i] = make([]bool, len(board[0]))
//	}
//	for i, v := range board {
//		for j := range v {
//			if dfs(board, visited, word, 0, i, j) {
//				return true
//			}
//		}
//	}
//	return false
//}
//
//func dfs(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
//	if index == len(word) - 1 {
//		return board[x][y] == word[index]
//	}
//	if board[x][y] == word[index] {
//		visited[x][y] = true
//		//开始在上下左右范围内搜索
//		for i := 0; i < 4; i++ {
//			nx := x + dir[i][0]
//			ny := y + dir[i][1]
//			if isInboard(board, nx, ny) && !visited[nx][ny] &&
//				dfs(board, visited, word, index + 1, nx, ny){
//				return true
//			}
//		}
//		visited[x][y] = false
//	}
//	return false
//}
////判断是否在矩阵中
//func isInboard(board [][]byte, nx int, ny int) bool {
//	return nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0])
//}

func exist(board [][]byte, word string) bool{
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++{
			if dfs(board, word, i, j, 0) {return true}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i int, j int, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {return false}
	if k == len(word) - 1 {return true}
	if board[i][j] == word[k]{
		temp := board[i][j]
		//标记已经被访问
		board[i][j] = ' '
		if dfs(board, word, i + 1, j, k + 1) || dfs(board, word, i - 1, j, k + 1) ||
			dfs(board, word, i, j + 1, k + 1) || dfs(board, word, i , j - 1, k + 1) {
			return true
		} else {
			//没有搜索成功，还原字符
			board[i][j] = temp
		}
	}
	return false
}

func main() {
	var a = ' '
	fmt.Printf("%#v", a)
}




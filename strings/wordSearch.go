package strings

import (
	"fmt"
)

func WordSearchInGrid(grid [][]string, word string ) {
	row := len(grid)
	col := len(grid[0])
	count := 0 
	for i := 0 ; i < row ; i ++ {
		for j := 0 ; j < col ; j ++ {
			if grid[i][j] == string(word[0]) && WordSearchInGridHelper(row, col, grid,word, count) {
				fmt.Println("TRUE")
			}
		}
	}
}


func WordSearchInGridHelper(row int , col int, board [][]string, word string, count int) bool {

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != string(word[count]) {
		return false
	}   

	if count == len(word) {
		return true
	}
	temp := board[row][col]
	board[row][col] = ""
	found :=  WordSearchInGridHelper(row+1,col,board,word,count+1) ||
	WordSearchInGridHelper(row,col+1,board,word,count+1) || 
	WordSearchInGridHelper(row-1,col,board,word,count+1) || 
	WordSearchInGridHelper(row,col-1,board,word,count+1) 
	board[row][col] = temp
	return found
	
}
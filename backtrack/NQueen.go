package backtrack

import (
	"fmt"
)

func NQueen(n int) {
	ans := [][]string{}
	board := make([][]string, n)
	for i := 0 ; i < n ; i ++ {
		board[i]=make([]string, n)
	}
	for i := 0 ; i < len(board[0]) ; i ++ {
		for j := 0 ; j < len(board) ; j++ {
			board[i][j] = "."
		}
	}
	NQueenHelper(board, &ans, 0, n)
	fmt.Println(board)
}

func NQueenHelper(board [][]string, ans *[][]string, col int, n int) {
	if col == n {
		return
	}
	for row := 0 ; row < n ; row ++ {
		if IsSafe(board, row, col, n) {
			board[row][col] = "Q"
			NQueenHelper(board,ans,col+1, n)
			board[row][col] = "."
		}
	}
}

func IsSafe(board [][]string, row,col int, n int) bool {
	dupRow := row
	dupCol := col 
	for row >= 0 && col >= 0 { //diagonal left to right
		if board[row][col] == "Q" {
			return false
		}
		row--
		col--
	}

	row = dupRow
	col = dupCol
	for col >=0 {
		if board[row][col] == "Q"{
			return false
		}
		col--
	}

	row = dupRow
	col = dupCol
	for row < n && col >=0 {
		if board[row][col] == "Q" {
			return false
		}
	}
	return true
}

func Test(n int){
	chess := make([][]string, n)
	for i := range chess {
		chess[i] = make([]string, n)
		for j := range chess[i] {
			chess[i][j] = "."
		}
	}
	fmt.Println(chess)
}

func NQueenTest(n int) {
	//ans := [][]string{}
	board := make([][]string, n)
	for i := 0 ; i < n ; i ++ {
		board[i]=make([]string, n)
	}
	for i := 0 ; i < len(board[0]) ; i ++ {
		for j := 0 ; j < len(board) ; j++ {
			board[i][j] = "."
		}
	}
	fmt.Println(board)
}
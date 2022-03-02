package backtrack

func RatInMaze(board [][]int, N int) {
	var result [][]string
	visitor := make([][]int, len(board))
	for i := 0 ; i < len(board) ; i ++ {
		visitor[i]=make([]int,len(board[0]))
	}
	di := []int{1,0,0,-1}
	dj := []int{0,-1,1,0}
	RatInMazeHelper(board,visitor,di,dj, &result)
}

func RatInMazeHelper(board [][]int, visitor[][]int, di, dj []int, result *[][]string){
	//start for loop with zero because for every positio you try DLRU
}
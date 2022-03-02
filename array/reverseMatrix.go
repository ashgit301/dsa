package array
import (
	"fmt"
)
func ReverseMatrix(arr [][]int) {
	//fmt.Println(arr)
	r := len(arr)-1
	c := len(arr[0])-1
	for i := 0 ; i < r ; i ++ {
		for j := 0 ; j < c ; j ++ {
			fmt.Println(arr[i][j], arr[j][i])
			Swap(arr[i][j], arr[j][i])
		}
	}
	for i, j := 0, r; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}
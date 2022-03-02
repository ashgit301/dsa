package array

import (
	"fmt"
)

func PascalTriange(n int){
	mx := make([][]int, n)
	for i := 0 ; i < n ; i ++ {
		mx[i]=make([]int, i)
	}
	for i :=0 ; i < n ; i++{
		for j :=0 ; j < len(mx[i]) ; j++ {
			if j == 0 || j == len(mx[i])-1 {
				mx[i][j] =1
			}else {
				mx[i][j] = mx[i-1][j] + mx[i-1][j-1]
			}
		}
	}
	fmt.Println(mx)
}
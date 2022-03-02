package array

import (
	"fmt"
)

func GridUniquePath(m,n int) {
	mx := make([][]int, m)
	for i := 0 ; i < len(mx) ; i ++ {
		mx[i]=make([]int, n)
	}
	fmt.Println(">>MX: ", mx)
	for j := 0 ; j < n ; j ++ {
		mx[0][j] = 1
	}
	for i := 0 ; i < m ; i ++ {
		mx[i][0] = 1
	}
	fmt.Println(">>MX: ", mx)
	for i := 1 ; i < m ; i ++ {
		for j := 1 ; j < n ; j ++ {
			mx[i][j] = mx[i-1][j] + mx[j-1][i]
		}
	}
	fmt.Println(mx[m-1][n-1])
}

func GridUniquePath2(m,n int) int {
	if m == 1 || n ==1 { //first row or first col, return 1
		return 1 
	}
	return GridUniquePath2(m-1,n) + GridUniquePath2(m,n-1)
}

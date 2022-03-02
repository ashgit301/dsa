package array

import (
	"fmt"
)

func SearchInMatrix(arr [][]int, k int) {
	r := len(arr)-1
	c := len(arr[0])-1
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j ++ {
			if arr[i][j] == k {
				//true
			}
		}
	}
}

func SearchInMatrix2(matrix [][]int, k int) {
	m := len(matrix)-1
	n := len(matrix[0])-1
	lo := 0
	hi := m*n
	for lo < hi {
		mid  := (lo + (hi - lo) / 2)
		if (matrix[mid/m][mid % m] == k){
			//yes
		}else if (matrix[mid/m][mid % m] > k){
			hi = mid-1
		}else {
			lo = mid + 1
		}
	}
} 
//if matrix is sorted
func SearchInMatrix3(arr [][]int, k int) {
	m := len(arr)-1
	n := len(arr[0])-1
	i := 0
	j := len(arr[0])-1
	for i>=0 && j >=0 && i<=  m && j <= n {
		if arr[i][j] == k {
			fmt.Println("True")
			break
		}else if arr[i][j] < k{
			i++
		}else {
			j--
		}
	}
	
}
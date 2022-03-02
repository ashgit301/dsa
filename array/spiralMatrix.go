package array

import (
	"fmt"
)

func SpiralMatrix(arr[][]int) {
	m := len(arr)
	n := len(arr[0])
	top := 0 
	left := 0
	bottom := m 
	right := n
	//pos := []int{0,1,2,3}
	dir := 0
	for top < m && bottom >=0 && left < m && right >=0 {
		if dir == 0 {
			for i := left ; i < right ; i ++ {
				fmt.Println(arr[top][i])
			}
			top++
		}else if dir == 1 {
			for i := top ; i < bottom ; i ++ {
				fmt.Println(arr[i][right])
			}
			right--
		}else if dir == 2 {
			for i := right ; i < left ; i -- {
				fmt.Println(arr[bottom][i])
			}
			bottom--
		}else {
			for i := bottom ; i < top ; i -- {
				fmt.Println(arr[bottom][i])
			}
			left--
		}
		dir = (dir+1)%4
	}
}
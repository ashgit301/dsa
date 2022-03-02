package array

import (
	"fmt"
)

func MissingAndRepeat(arr []int) {
	countArr := make([]int, 5)
	for i := 0 ; i < len(arr) ; i ++ {
		countArr[i]++
	}
	for i := 0 ; i < len(countArr) ; i ++ {
		if countArr[i] == 0 || countArr[i] > 1 {
			fmt.Println("Missing & Repeat", arr[i])
		}
	}
}
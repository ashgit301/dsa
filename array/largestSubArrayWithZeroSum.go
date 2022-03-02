package array

import (
	"fmt"
)


func SubArrayZero(arr []int) {
	maxStreak := 0  
	for i := 0 ; i < len(arr) ; i ++ {
		sum := 0 
		for j := i; j < len(arr) ; j ++ {
			sum = sum + arr[j]
			if sum == 0 {
				maxStreak = Max(maxStreak, j - (i+1))
			}
		}
	}
	fmt.Println(maxStreak)
}
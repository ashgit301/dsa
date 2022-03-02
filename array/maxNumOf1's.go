package array

import (
	"fmt"
)


func MaxOnes(arr []int) {
	currStreak := 0 
	maxStreak := 0
	for i := 0 ; i < len(arr) ; i ++ {
		if arr[i] == 1 {
			currStreak++
			maxStreak = Max(maxStreak,currStreak)
		}else {
			currStreak = 0
		}
	}
	fmt.Println(maxStreak)
}
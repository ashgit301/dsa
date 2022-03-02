package array

import (
	"fmt"
)

func JumpGame(arr []int) {
	pos := 0 
	ans := helperJump(pos,arr, 0)
	fmt.Println(ans)
}

func helperJump(pos int, arr []int, count int) bool {
	fmt.Println("position :", pos)
	if pos >= len(arr)-1 {
		return true
	}
	farthest := Min(pos + arr[pos], len(arr)-1)
	for i := pos ; i <= farthest ; i ++ {
		if (helperJump(i+1,arr, count+1)){
			return true
		}
	}
	return false
	
}

func HelperJump2(arr []int) {
	goal := len(arr)-1
	for i := len(arr)-2 ; i >= 0 ; i -- {
		if arr[i] + i >= goal {
			goal = i 
		}
	}
	if goal == 0 {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
}

func Reachable(arr []int) bool {
	Reachable := 0 
	for i := 0 ; i < len(arr) - 1 ; i ++ {
		if Reachable < i {
			return false
		}
		Reachable = Max(Reachable,i+arr[i])
	}
	return true
}

func MinNumJumps(arr []int) {
	jump := 1
	ladder := arr[0]
	stairs := arr[0]
	for i := 1 ; i < len(arr)-1 ; i ++ {
		if arr[i] + i > ladder {
			ladder = arr[i] + i
		}
		if stairs == 0 {
			jump++
			stairs = ladder-i
		}
		stairs--
	}
	fmt.Println(jump)
}
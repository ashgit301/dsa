package array

import (
	"fmt"
)

func MaximumSubArray(arr []int) {
	sum := 0 
	maxi := 0 
	for i := 0 ; i < len(arr) ; i ++ {
		sum = 0 
		for j := i ; j < len(arr) ; j ++ {
			sum = sum + arr[j]
			maxi = Max(sum,maxi)
		}
	}
	fmt.Println(maxi)
}


func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func Kadane(arr []int) {
	maxSoFar := arr[0]
	maxEndHere := arr[0]
	for i := 1 ; i < len(arr) ; i ++ {
		maxEndHere = Max(maxEndHere+arr[i], maxEndHere)
		maxSoFar = Max(maxSoFar, maxEndHere)
	}
	fmt.Println(maxSoFar)
}

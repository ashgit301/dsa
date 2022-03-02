package array

import (
	"fmt"
)

func Trapping(arr []int) {
	res := 0 
	for i := 1 ; i < len(arr)-1 ; i ++ { //ignore 1st & last
		leftMax := 0 
		for j := 0 ; j < i ; j ++ {
			leftMax = Max(leftMax, arr[j])
		}
		rightMax := 0 
		for k := i+1 ; k < len(arr); k ++ {
			rightMax = Max(rightMax, arr[k])
		}
		fmt.Println(arr[i],leftMax,rightMax)
		res = res + (Min(leftMax, rightMax) - arr[i])
	}
	fmt.Println(res)
}

//simlarly form prefix & suffix array also with this method

func Trapping2(arr []int) {
	leftMax := 0 
	rightMax := 0 
	res := 0 
	lo := 0 
	hi := len(arr)-1 
	for lo <= hi {
		if leftMax <= rightMax {
			if leftMax < arr[lo] {
				leftMax = arr[lo]
			}else {
				res = res+leftMax-arr[lo]
			}
			lo++
		}
		if leftMax >= rightMax {
			if rightMax < arr[hi] {
				rightMax = arr[hi]
			}else {
				res = res+rightMax-arr[hi]
			}
			hi--
		}
	}
	fmt.Println(res)
}
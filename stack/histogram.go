package stack

import (
	"fmt"
)
 
func Histogram(arr []int) {
	maxArea := 0 
	for i := 0 ; i < len(arr) ; i ++ {
		leftSmall := 0 
		rightSmall := len(arr)-1
		for j := i+1 ; j < len(arr) ; j ++ {
			if arr[i] > arr[j] {
				rightSmall = j-1
				fmt.Println("index: ", i, "rightSmall : ", rightSmall)
				break
			}
 	    }
		for k := i ; k > 0 ; k -- {
			if arr[i] > arr[k] {
				leftSmall = k+1
				fmt.Println("index: ", i, "leftSmall : ", leftSmall)
				break
			}
 		}
		maxArea = Max(maxArea, (rightSmall - leftSmall + 1 )*arr[i])	
	}
	fmt.Println(maxArea)
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}


func HistogramDynamicPrograming(arr []int) {
	max_area := 0 
	for i := 0 ; i < len(arr) ; i ++ {
		min_element := 100000
		for j := i ; j < len(arr) ; j ++ {
			min_element = Min(min_element, arr[j])
			max_area = Max(max_area,min_element*(j-i+1))
		}
	}
}


func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}


func HistogramStack(arr []int) {
	stack := make([]int, len(arr))
	leftSmall := make([]int, len(arr))
	//rightSmall := make([]int, len(arr))
	for i := 0 ; i < len(arr) ; i ++ {
		if len(stack) != 0 {
			for st := len(stack)-1 ; st >= 0 ; st -- {
				if arr[stack[st]] > arr[i] {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if len(stack) == 0 {
			leftSmall[i] = 0  
		}else {
			leftSmall[i] = stack[len(stack)-1]
		}
		stack = append(stack,i)
	}
	fmt.Println(leftSmall)
}



package array

import (
	"fmt"
)

func FindPivot(arr []int) {
	for i := 0 ; i < len(arr) ; i ++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr[i+1])
		}
	}
}

func FindPivot2(arr []int) {
	low := 0 
	high := len(arr)
	for low <= high {
		mid := low+high/2
		if arr[mid] > arr[mid+1] {
			fmt.Println(arr[mid])
		}else if arr[low] <= arr[mid] { //right sorted
			low = mid + 1  
		}else {
			high = mid - 1
		}
	}
}
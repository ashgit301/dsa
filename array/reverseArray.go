package array

import (
	"fmt"
)

func ReverseArray(arr []int) {
	lo := 0 
	hi := len(arr)-1
	for lo < hi  {
		arr[lo] , arr[hi] = arr[hi] , arr[lo]
		lo++
		hi--
	}
	fmt.Println(arr)
}
package array

import (
	"fmt"
)

func PositiveNegativeToSide(arr []int) {
	l := 0
	r := len(arr)-1
	for l < r {
		if arr[l] < 0 {
			l++
		}else if arr[l] > 0 && arr[r] < 0 {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}else {
			r--
		}
 	}
	fmt.Println(arr)
}

func PositiveNegativeToSide2(arr []int) {
	j := 0 
	for i := 0 ; i < len(arr) ; i++{
		if arr[i] < 0 {
			if i != j {
				arr[i],arr[j] = arr[j],arr[i]
			}
			j++
		}
	}
	fmt.Println(arr)
}
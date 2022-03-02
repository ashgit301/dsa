package array

import (
	"fmt"
)

func InsertionSort(arr []int){
	for i := 1 ; i < len(arr) ; i++ {
		key := arr[i]
		hole := i 
		for hole > 0 && arr[hole-1] > arr[hole] {
			arr[hole] = arr[hole-1]
			hole-- 
		}
		arr[hole] = key
	}
	fmt.Println(arr)
}

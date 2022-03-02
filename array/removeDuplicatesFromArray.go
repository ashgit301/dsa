package array

import (
	"fmt"
)

func RemoveDuplicates(arr []int) {
	mpp := map[int]int{}
	count := 0
	for i := 0 ; i < len(arr) ; i ++ {
		if _, ok := mpp[arr[i]]; !ok {
			mpp[arr[i]]++
		}
	}
	for k,_ := range mpp {
		_ = k
		count ++
	}
	fmt.Println(count)
}

func RemoveDuplicates2(arr []int ) {
	count := 0
	i := 0 
	for j := 1 ; j < len(arr) ; j ++ {
		if arr[i] != arr [j] {
			count++
			i = j 
		}
	}
	fmt.Println(count+1)
}
package array

import (
	"fmt"
)

func MajorityElement(arr []int) {
	length := len(arr)/2
	mpp := map[int]int{}
	for i := 0 ; i < len(arr) ; i ++ {
		mpp[arr[i]]++
	}
	fmt.Println("Map >> ", mpp)
	for key,val := range mpp {
		if val > length{
			fmt.Println(key)
		}
	}
}

func Moores(arr []int) {
	count := 0
	candidate := 0 
	for i := 0 ; i < len(arr) ; i ++ {
		if count == 0 {
			candidate = arr[i]
		}
		if arr[i] == candidate {
			count++
		}else {
			count--
		}
	}
}

//intuiton is same for n/3 also 
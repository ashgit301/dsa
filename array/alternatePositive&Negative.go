package array

import (
	"fmt"
)

func AlternatePositive_Negative(arr []int) { //incomplete
	res := make([]int, len(arr))
	pos := []int{}
	neg := []int{}
	for i := 0 ; i < len(arr) ; i ++ {
		if arr[i] >=0 {
			pos = append(pos, arr[i])
		}else {
			neg = append(neg, arr[i])
		}
	}
	m := 0 
	n := 0 
	for i := 0 ; i < len(res) ; i ++  {
		if i % 2 == 0 {
			res[i] = pos[m]
			m++
		}else {
			res[i] = neg[n]
			n--
		}
	}
}

func AlternatePositive2(arr []int) {
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
	neg := 0
	pos := 0
	for arr[pos] < 0 {
		pos++
	}
	fmt.Println(arr[neg])
	fmt.Println(arr[pos])
	for pos <=len(arr)-1 && neg <= len(arr)-1 {
		if pos > neg {
			arr[neg], arr[pos] = arr[pos], arr[neg]
			neg = neg + 2 
			pos = pos + 1
		}else {
			break
		}
	}
	fmt.Println(arr)

}
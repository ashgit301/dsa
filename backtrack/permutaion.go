package backtrack

import (
	"fmt"
)

func Permutaion(arr []int) {
	var result [][]int 
	visitor := make([]bool, len(arr))
	PermutaionHelper(arr, &result, []int{}, visitor)
	fmt.Println(result)
}

func PermutaionHelper(arr []int , result *[][]int, temp[]int, visitor []bool){
	if len(temp) == len(arr) {
		t := make([]int, len(temp))
		copy(t, temp)
		*result = append(*result, temp)
		return
	}

	for i := 0 ; i < len(arr) ; i ++ {
		if visitor[i] == false {
			visitor[i] = true
			temp = append(temp, arr[i])
			PermutaionHelper(arr, result, temp ,visitor)
			temp = temp[:len(temp)-1]
			visitor[i] = false
		}
	}
	return
}
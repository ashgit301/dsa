package backtrack

import (
	"fmt"
)

func CombinationSum(arr []int, k int) {
	var result [][]int 
	CombinationSumHelper(arr, &result, []int{}, 0, 7)
	fmt.Println(">> ", result)
}

func CombinationSumHelper(arr []int, result *[][]int, temp []int, start int, target int) {
	if target == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		*result = append(*result, temp)
		return
	}else if target < 0 {
		return
	}
	for i := start ; i < len(arr) ; i ++ {
		temp = append(temp, arr[i])
		CombinationSumHelper(arr, result,temp,start,target-arr[i])
		temp = temp[:len(temp)-1]
	}
	
}
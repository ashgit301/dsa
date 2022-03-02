package backtrack

import (
	"fmt"
)

func Subsetsum(arr[]int){
	var result []int
	ans := SubsetsumHelper(arr, 0,0, &result)
	fmt.Println(*ans)
}

func SubsetsumHelper(arr[]int, index int, sum int , result *[]int) *[]int{
	if index == len(arr){
		*result = append(*result, sum)
		return result
	}
	SubsetsumHelper(arr, index+1, sum+arr[index],result)
	SubsetsumHelper(arr, index+1, sum,result )
	return result
}

func Subsetsum2(nums []int) {
	var result []int
	ans := SubsetsumHelper2(nums, &result, 0, 0)
	fmt.Println(ans)
}

func SubsetsumHelper2(arr []int, result *[]int, index int, sum int) *[]int{
	fmt.Println(index,sum)
	*result = append(*result, sum)
	for i := index; i < len(arr); i++ {
		sum = sum+arr[i]
		SubsetsumHelper2(arr, result, i+1, sum)
		sum = sum-arr[i]
	}
	return result
}



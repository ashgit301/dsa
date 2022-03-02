package backtrack

import (
	"fmt"
)
func GenerateSubArray(arr []int) {
	result := [][]int{}
	for i := 0 ; i < len(arr) ; i ++ {
		curr := []int{}
		for j := i ; j < len(arr) ; j ++ {
			if j == len(arr)-1 {
				curr = append(arr[i:])
			}
			curr = append(arr[i:j+1])
			result = append(result,curr)
		}
	}
	fmt.Println(result)
}


func Subsets(nums []int) {
	var result [][]int
	helper(nums, &result, []int{}, 0)
	fmt.Println(result)
}

func helper(nums []int, result *[][]int, temp []int, index int) {
	fmt.Println(index, temp)
	t := make([]int, len(temp))
	copy(t, temp)
 	*result = append(*result, temp)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(nums, result, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}
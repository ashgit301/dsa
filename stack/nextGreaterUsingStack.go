package stack

import (
	"fmt"
)

func NextGreater(nums []int) {
	res := make([]int, len(nums))
	stack := []int{}

	for i := len(nums)-1; i >= 0 ; i -- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			res[i] = stack[len(stack)-1]
		}else {
			res[i] = -1
		}
		stack = append(stack, nums[i])
	}
	fmt.Println(res)
}
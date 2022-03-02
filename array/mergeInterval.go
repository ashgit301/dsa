package array

import (
	"fmt"
	"sort"
)

func MergerIntervals(nums [][]int){
	sort.Slice(nums, func(i, j int) bool {
        return nums[i][0] < nums[j][0]
    })
	res := [][]int{}
	curr := nums[0]
	if len(nums) == 1 {
		//return
	}
	for i := 1; i < len(nums) ; i++{
		if curr[1] >= nums[i][0]{
			curr[1] = Max(curr[1],nums[i][0])
		}else {
			res = append(res, curr)
			curr = nums[i]
		}
	}
	res = append(res, curr)
	fmt.Println(res)
}
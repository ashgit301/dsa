package array

import (
	"fmt"
)

func findDuplicate(nums []int) {
    if len(nums) == 1 {
       
    }
    slow := nums[0]
    fast := nums[nums[0]]
    for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	fmt.Println(nums)
}
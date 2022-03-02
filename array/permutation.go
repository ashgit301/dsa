package array

import (
	"fmt"
)

func Permutation(num []int) {
	freq := make([]bool, len(num))
	ds := []int{}
	ans := [][]int{}
	permutationHelper(freq,ds,num, &ans)
	fmt.Println(ans)
}

func permutationHelper(freq []bool, ds []int, nums []int, ans *[][]int){
	if len(ds) == len(nums) {
		*ans = append(*ans, ds)
	}
	for i := 0 ; i < len(nums) ; i++ {
		if !freq[i]{
			fmt.Println("index :", "frq : ",i, freq, ans )
			freq[i] = true
			ds = append(ds,nums[i])
			permutationHelper(freq,ds,nums,ans)
			freq[i]=false
			ds = ds[:len(ds)-1]
		}
	}
}
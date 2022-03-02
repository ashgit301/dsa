package array

import (
	"fmt"
)
//cant get the intuition behind this
func MinMaxDiff(arr []int, k int) {
	n := len(arr)
	min := arr[0]
	max := arr[len(arr)-1]
	ans := max - min
	for i := 1 ; i < len(arr) ; i++ {
		min = Min(arr[0]+k, arr[i]-k)
		max = Max(arr[n-1]-k,arr[i-1]+k)
		ans = Min(ans, max-min)
	}
	fmt.Println(ans)
}
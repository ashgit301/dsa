package array

import (
	"fmt"
	"sort"
)

func Max_And_Min(arr []int) {
	sort.Ints(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[len(arr)-1])
}

type mm struct {
	Min int
	Max int
}

func Max_And_Min_2(arr []int) {
	ans := mm{}
	if len(arr) == 1 {
		ans.Min = arr[0]
		ans.Max = arr[0]
	}

	if len(arr) > 1 {
		if arr[0] > arr[1] {
			ans.Max = arr[0]
			ans.Min = arr[1]
		}else {
			ans.Max = arr[1]
			ans.Min = arr[0]
		}
			
	}

	for i := 2 ; i < len(arr) ; i ++ {
		if arr[i] > ans.Max {
			ans.Max = arr[i]
		}
		if arr[i] < ans.Min {
			ans.Min = arr[i]
		}
	}
}


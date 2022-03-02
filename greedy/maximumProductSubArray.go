package greedy

import (
	"fmt"
)

func MaximumSubArrayProduct(arr []int) {
	maxProd := arr[0]
	for i := 1 ; i < len(arr) ; i ++ {
		prod := 1
		for j := i ; j < len(arr) ; j ++ {
			prod = prod * arr[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	fmt.Println(maxProd)
}

//can use kadanes algorithm also

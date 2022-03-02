package array

import (
	"fmt"
)

func MedianFromSortedArray(arr1,arr2 []int) {
	//can also make an array of size a+b and store the array using two pointers in sorted manner
	if len(arr2) < len(arr1) {
		MedianFromSortedArray(arr2,arr1)
	}
	left1 := 0 
	left2 := 0
	right1 := 0
	right2 := 0
	n1 := len(arr1)
	n2 := len(arr2)
	low := 0 
	high := n1
	for low <= high {
		cut1 := low+high/2 //median
		cut2 := ((n1+n2+1)/2)-cut1
		if cut1 == 0 {
			left1 = -1000
		}else {
			left1 = arr1[cut1-1]
		}
		if cut2 == 0 {
			left2 = -1000
		}else {
			left2 = arr2[cut2-1]
		}
		if cut1 == n1 {
			right1 = 1000
		}else {
			right1 = arr1[cut1]
		}
		if cut2 == n2 {
			right2 = 1000
		}else {
			right2 = arr2[cut2]
		}
		if left1 <= right2 && left2 <= right1 {
			if n1+n2 % 2 == 0 {
				fmt.Println((Max(left1,left2) + Min(right1, right2))/2.0)
			}else {
				fmt.Println(Max(left1,left2))
			}
			break
		}else if left1 > right2 {
			high = cut1-1
		}else {
			low = cut1+1
		}
	}

}
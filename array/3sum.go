package array

import (
	"sort"
	"fmt"
)

func Three_Sum(arr []int) {	//gives repeated elements // see strivers answer for this issue - unhash
	res := [][]int{}
	mpp := map[int]int{}
	for i := 0 ; i < len(arr) ; i ++ {
		mpp[arr[i]] ++ 
	}
	fmt.Println(mpp)

	for i := 0 ; i < len(arr) ; i ++ {
		for j := i+1 ; j < len(arr) ; j ++ {
			comp := arr[i]+arr[j]
			if Contains(mpp,-comp){
				curr := []int{arr[i], arr[j], -comp}
				res = append(res, curr)
			}
		}
	}
	fmt.Println(res)
}

func Three_sum2(arr []int) {
	sort.Ints(arr)
	res := [][]int{}
	for i := 0 ; i < len(arr)-2 ; i ++ {
		if i == 0 || i > 0 && arr[i-1] != arr[i]{
			lo := i+1
			hi := len(arr) -1 
			comp := 0 - arr[i]
			for lo < hi {
				if comp == arr[lo] + arr[hi] {
					curr := []int{arr[lo], arr[hi], -comp}
					res = append(res, curr)
					for lo < hi {
						if arr[lo] == arr[lo+1] {
							lo++
						}else if arr[hi] == arr[hi-1]{
							hi--
						}else {
							break
						}
				    }
					lo++
					hi--
				}else if (arr[lo] + arr[hi] < comp){
					lo++
				}else{ 
					hi--
				}
			}
	    }
	}
	fmt.Println(res)
}
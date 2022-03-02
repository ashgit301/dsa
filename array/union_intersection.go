package array

import (
	"fmt"
)

func Union_Intersection(arr1, arr2 []int) {
	res := []int{}
	i := 0 
	j := 0 
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		}else if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		}else {
			res = append(res, arr1[i])
			i++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	fmt.Println(res)
}

func Intersection(arr1, arr2 []int) {
	res := []int{}
	i := 0 
	j := 0 
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			i++
		}else if arr1[i] > arr2[j] {
			j++
		}else {
			//to check for duplicate see if it is already present in the res
			res = append(res, arr1[i])
			i++
		}
	}
}
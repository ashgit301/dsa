package array

import (
	"fmt"
)


func CyclicallyRotate(arr []int) { //wrong solution // read prob statement properly
	res := make([]int, len(arr))
	key1 := arr[0]
	key2 := arr[len(arr)-1]
	for i := 1; i < len(arr)-1 ; i++ {
		res[i] = arr[i]
	}
	res[0] = key2
	res[len(res)-1] = key1
	fmt.Println(res)
}

func CyclicallyRotate2(arr []int) {
	key2 := arr[len(arr)-1]
	for i := len(arr)-1 ; i > 0 ; i -- {
		arr[i] = arr[i-1]
	}
	arr[0] = key2
	fmt.Println(arr)
}
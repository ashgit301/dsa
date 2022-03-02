package array
import (
	"fmt"
)
func SingleElement(arr []int) {
	mpp := map[int]int{}
	for i := 0 ; i < len(arr) ; i ++ {
		mpp[arr[i]]++
	}
	fmt.Println(mpp)
	for key, _ := range mpp {
		//fmt.Println(mpp[key])
		if mpp[key] == 1 {
			fmt.Println(key)
		}
	}

}

func SingleElement2(arr []int) { //can be easily solves udig xor
	low := 0                     // |
 	high := len(arr)-2 // cus 113344|5
	for low <= high {             //|
		mid := high+low / 2 
		if mid % 2 == 0 {
			//right half second index
			//left half first index
			if arr[mid] == arr[mid-1] {
				high = mid-1 //right
			}else {
				low = mid+1 //left
			}
		}else {
			//right half first index
			//left half second index
			if arr[mid] == arr[mid+1] {
				high = mid-1//right
			}else {
				low = mid+1
			}
		}
	}
	fmt.Println(arr[low])
}
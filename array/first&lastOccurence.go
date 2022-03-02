package array
import("fmt")
func FindFirst_LastOccurace(arr []int, k int) {
	first := -1 
	last := -1
	for i := 0 ; i < len(arr) ; i ++ {
		if arr[i] == k {
			if first == -1 {
				first = i 
			}else {
				last = i
			}
		}
	}
	fmt.Println(first,last)
}

func FindFirst_LastOccurace2(arr []int, key int) { //same for last occurance
	first := -1
	low := 0 
	high := len(arr)-1
	for low <= high {
		mid := (low+high)/2
		fmt.Println(">> ", low,high,mid)
		if arr[mid] == key && arr[mid-1] < key {
			first = mid
			break
		}else if arr[mid] < key {
			low = mid+1
		}else {
			high = mid-1 
		}
	}
	fmt.Println(first)
}
package array

func RotatedSortedArray(arr[]int, key int) {
	low := 0 
	high := len(arr)-1 
	for low <= high {
		mid := low+high/2 
		if arr[low] <= arr[mid] { // if left is sorted
			if key >= arr[low] && key <= arr[mid] { // id lies in left part
				high = high - 1
			}else {
				low = mid + 1 
			}
		}else {
			if key >= arr[mid] && key <= arr[high] { //if right is sorted
				low = mid +1 
			}else {
				high = mid - 1
			}
		}
	}
}
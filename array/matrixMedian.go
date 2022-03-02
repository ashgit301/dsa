package array
//loop through the matrix & store the elements in an array. find the median.
//optimal solution
func MatrixMedian(arr [][]int) {
	low := 0 
	high := 999999
	m := len(arr)
	n := len(arr[0])
	for low <= high {
		count := 0
		mid := low + high / 2
		for i := 0 ; i < m ; i ++ {
			count = count + MinAvailable(mid,arr[i])
		}
		if count <= n*m/2 {
			low = mid+1 
		}else {
			high = mid -1 
		}
	}
	//return l
}

func MinAvailable(key int,arr[]int ) int {
	low := 0 
	high := len(arr)-1
	for low <= high {
		mid := low+high/2
		if arr[mid] > key {
			low = mid+1
		}else {
			high = mid-1
		}
	}
	return low //dont know why??
	//cus we are finding index where all elements before it is less than the key
	
}
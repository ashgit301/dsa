package array

func NextPermutation(arr []int) {
	n := len(arr)
	firstIndex := 0 
	for i := n-2 ; i >=0 ; i ++ { //findign the break point 
		if arr[i] < arr[i+1] {
			firstIndex = i
			break
		}
	}
	if firstIndex < 0 {
		Reverse(arr,0, n-1)
	}else { //searh for num greater than the break point 
	    j := n-1
		for j >= 0 && arr[j] <= arr[firstIndex] {
			j--
		}
		arr[firstIndex], arr[j] = arr[j], arr[firstIndex]
		Reverse(arr, firstIndex+1, len(arr)-1)
	}
}





func Reverse(arr []int, start, end int) {
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++; end--
    }
}
package array

//import("fmt")

func QuickSort(arr []int, low, high int) {
	//fmt.Println(arr)
	if low < high {
		var pivot = partition(arr, low, high)
		QuickSort(arr, low, pivot)
		QuickSort(arr, pivot + 1, high)
	}
}

func partition(arr []int, low, high int) int {
	var pivot = arr[low]
	var i = low
	var j = high

	for i < j { //this line is the key point to note
		for arr[i] <= pivot && i < high {
			i++;
		}
		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot

	return j
}
package array

func Merge(Arr []int, l int, r int, mid int) {
	i := l
	j := r 
	k := l 
	tmp := []int{}
	for i <=l && j <= r {
		if Arr[i] < Arr[j] {
			tmp[k] = Arr[i]
			k++
			i++
		}else {
			tmp[k] = Arr[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmp[k] = Arr[i]
	}
	for j <= r {
		tmp[k] = Arr[j]
	}
}

func RecursiveMergeSort(Arr []int, low int, high int) {
	if (low < high){
        // Calculate mid point
        mid := low + (high-low)/2;
        // Sort sub-lists
        RecursiveMergeSort(Arr, low, mid);
        RecursiveMergeSort(Arr, mid+1, high);
 
        // Merge sorted sub-lists
        Merge(Arr, low, mid, high);
    }
}
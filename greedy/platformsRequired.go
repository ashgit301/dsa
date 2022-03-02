package greedy

func PlatformsRequired(arr, dep []int) {
	result := 1 
	platform := 1 
	i := 0 
	j := 0 
	for i < len(arr) && j < len(dep) {
		if arr[i] <= dep[j] {
			i++
			platform++
		}else if arr[i] > arr[j] {
			j++
			platform--
		}
		if platform > result {
			result= platform
		}
	}
}
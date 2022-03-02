package array


func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for idx, v := range nums {
		comp := target - v
		if item, ok := result[comp]; ok {
			return []int{item, idx}
		}
		result[v] = idx
	}
	return []int{-1, 1}
} 
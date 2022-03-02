package array

func NthRoot(n,m int) {
	low := 1 
	high := m 
	e := 1e-6
	for high-low > int(e)  {
		mid := high+low /2
		comp := Multiply(mid, n)
		if high > comp {
			low = mid
		}else {
			high = mid
		}
	}
}

func Multiply(a,b int) int {
	ans := a
	for i := 1 ; i <= b ; i ++ {
		ans = ans * b
	}
	return ans
}
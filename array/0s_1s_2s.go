package array

func DutchNationalFlag(arr []int) {
	lo := 0
	mid := 0 
	hi := len(arr)-1
	for mid <= hi {
		if arr[mid] == 0{
			Swap(lo,mid)
			lo++
			mid++
		}else if arr[mid]==1{
			mid++
		}else {
			Swap(mid,hi)
			hi--
		}
	}

}

func Swap(a,b int){
	tmp := a 
	a = b 
	b = tmp
}
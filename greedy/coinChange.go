package greedy

func CoinChange(N int) {
	deno := []int{1,2,5,10,20,50,100,500,1000}
	res := 0 
	want := N
	for i := len(deno)-1 ; i >=0 ; i ++ {
		for deno[i] <= N {
			res++
			want = want-deno[i]
		}
	}
}
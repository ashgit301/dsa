package array

import (
	"fmt"
)

func StockBuyAndSell(arr []int){ //idk why wrong solution
	fmt.Println("Stock")
	profit := 0
	for i := 0 ; i < len(arr) ; i ++ {
		for j := i+1 ; j < len(arr) ; j++ {
			if arr[i] < arr[j] {
				fmt.Println(arr[j]-arr[i])
				profit = Max(profit, arr[i]-arr[j])
			}
		}
	}
}

func StockBuyAndSell2(arr []int) {
	min := 10000000
	profit := 0
	for i := 0; i < len(arr) ; i ++ {
		min = Min(min,arr[i])
		profit = Max(profit, arr[i]-min)
	}
	fmt.Println(profit)
}


func Min(a int, b int)int {
	if a > b {
		return b
	}else if b > a {
		return a 
	}else {
		return a
	}
}
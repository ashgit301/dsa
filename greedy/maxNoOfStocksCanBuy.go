package greedy

import (
	"fmt"
	"sort"
)

type Stocks struct {
	Price int
	Num int
}

func MaxNumOfStocks(stocks []int, k int) {
	maxCount := 0	
	total := 0 
	arrOfStocks := make([]Stocks, len(stocks))
	for i := 0 ; i < len(arrOfStocks) ; i ++ {
		arrOfStocks[i].Price = stocks[i]
		arrOfStocks[i].Num = i+1
	}
	sort.Slice(arrOfStocks, func(i,j int)bool {
		return arrOfStocks[i].Price < arrOfStocks[j].Price
	})
	fmt.Println(arrOfStocks)
	for i:= 0 ; i < len(arrOfStocks) ; i ++ {
		for j := arrOfStocks[i].Num ; j > 0 ; j -- {
			total = total + arrOfStocks[i].Price
			if total <= k {
				maxCount++
			}
		}
	}
	fmt.Println(maxCount)
}
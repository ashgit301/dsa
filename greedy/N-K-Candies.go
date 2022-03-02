package greedy

import (
	"fmt"
	"sort"
)

func NnumOfCandy(price []int, k int) {
	cost := 0
	sort.Ints(price)
	numOfCandy := len(price)
	for i := 0 ; i < numOfCandy-k ; i ++ {
		cost = cost+price[i]
	}
	fmt.Println(cost)
}

func NnumOfCandyMax(price []int, k int) {
	cost := 0 
	sort.Slice(price, func(i,j int)bool {
		return price[i] > price[j]
	})
	fmt.Println(price)
	numOfCandy := len(price)
	for i := 0 ; i < numOfCandy-k ; i ++ {
		cost = cost+price[i]
	}
	fmt.Println(cost)


}

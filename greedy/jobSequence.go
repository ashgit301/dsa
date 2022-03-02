package greedy

import (
	"fmt"
	"sort"
)
type Job struct {
	ID int
	DeadLine int
	Profit int
}
func JobSequence(deadLine, profit []int) {
	sortedJob := make([]Job, len(deadLine))
	for i := 0 ; i < len(deadLine) ; i ++ {
		sortedJob[i].ID = i 
		sortedJob[i].DeadLine = deadLine[i]
		sortedJob[i].Profit = profit[i]
	}
	sort.Slice(sortedJob, func (i,j int)bool{
		return sortedJob[i].Profit > sortedJob[j].Profit
	})
	datesAvailable := make([]int, (sortedJob[len(sortedJob)].Profit+1))
	MaxProfit := 0
	jobCount := 0
	for i := 0 ; i < len(sortedJob) ; i ++ {
		for j := sortedJob[i].DeadLine ; j > 0 ; j-- {
			if datesAvailable[j] == 0 {
				datesAvailable[j] = -1
				MaxProfit = MaxProfit+sortedJob[i].Profit
				jobCount++
				break
			}
		} 
	}
	fmt.Println(MaxProfit)
}
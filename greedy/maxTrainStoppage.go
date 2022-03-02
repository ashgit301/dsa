package greedy

import (
	"sort"
	"fmt"
)


type TrainStoppage struct {
	Arrival int
	Departure int
	Platform int
}

func MaxTrainStoppage(arr [][]int, n int) {
	maxTrain := 0
	platformArr:= make([]int, n+1)
	trainArr := make([]TrainStoppage, len(arr))
	fmt.Println(trainArr)
	for i := 0 ; i < len(trainArr) ; i ++ {
		trainArr[i].Arrival = arr[i][0]
		trainArr[i].Departure = arr[i][1]
		trainArr[i].Platform = arr[i][2]
	}
	sort.Slice(trainArr, func(i,j int)bool{
		return trainArr[i].Departure < trainArr[j].Departure //if want smaller point to j
	})
	for i := 0 ; i < len(trainArr) ; i ++ {
		if platformArr[trainArr[i].Platform] == 0 {
			maxTrain++
			platformArr[trainArr[i].Platform] = trainArr[i].Departure
		}else if platformArr[trainArr[i].Platform] < trainArr[i].Arrival {
			maxTrain++
			platformArr[trainArr[i].Platform] = trainArr[i].Departure
		}
	}
	fmt.Println(maxTrain)
	
}
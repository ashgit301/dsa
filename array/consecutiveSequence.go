package array
import (
	"fmt"
	"sort"
)
func ConsecutinveSequence(arr []int) { //wring answer, wrong approach
	sort.Ints(arr)
	fmt.Println(arr)
	Maxcount := 0 
	count := 0 
	for i := 1 ; i < len(arr); i ++ {
		comp := arr[i-1]
		if arr[i] - comp == 1 {
			count ++
			Maxcount = Max(Maxcount, count)
		}else {
			count = 0
		}

	}
	fmt.Println(Maxcount)
}

func ConsecutinveSequence2(arr []int) {
	mpp := map[int]int{}
	longestStreak := 0 
	for i := 0 ; i < len(arr) ; i ++ {
		mpp[i] = arr[i]
	}
	fmt.Println(mpp)

	for i := 0 ;  i < len(arr) ; i ++ {
		currStreak := 0 
		if ok := Contains(mpp, arr[i]-1); !ok {
			currStreak++
			comp := arr[i]
			for Contains(mpp, comp+1) {
					currStreak++
					comp = comp+1
					longestStreak = Max(longestStreak, currStreak)
			
			 	}
			}
		}
		fmt.Println(longestStreak)
	}

func Contains(m map[int]int, v int) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}
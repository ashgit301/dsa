package greedy

import (
	"fmt"
)

func ChooseAndSwap(s string){
	mpp := make(map[int]int, 123)
	flag := false
	for i := 0 ; i < 123 ; i++{
		mpp[i] = -1
	}
	for i := 0 ; i < len(s) ; i ++ {
		ascii := int(s[i])
		if mpp[ascii] == -1 {
			mpp[ascii] = i 
		}
	}
	fmt.Println(mpp)
	for i := 0 ; i < len(s) ; i ++ {
		currAscii := int(s[i])
		fmt.Println(currAscii)
		for j := 0 ; j < currAscii ; j ++ {
			if mpp[j] != -1 {
				//go the index 
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
}



package strings

import (
	"fmt"
)

func FindDuplicate(s string) {
	mpp := map[string]int{}
	for i := 0 ; i < len(s) ; i ++ {
		mpp[string(s[i])]++
	}
	fmt.Println(mpp)
	for key,_ := range mpp {
		if mpp[key] > 1 {
			fmt.Println("Repeated element : ", key)
		}
	}
}
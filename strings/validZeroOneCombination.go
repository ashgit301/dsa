package strings

import (
	"fmt"
)

func ZerosAndOnes(s string) {
	count0 := 0
	count1 := 1 
	count := 0 
	for i := 0 ; i < len(s) ; i ++ {
		if s[i] == '0'{
			count0++
		}else {
			count1++
		}
		if count0 == count1 {
			count++
		}

	}
	fmt.Println(count)
}
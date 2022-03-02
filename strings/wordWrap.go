package strings

import (
	"fmt"
	"strings"
)

func WordWrap(s string, k int) {
	str := strings.Fields(s)
	fmt.Println(str)
	freeSpace := k
	cost := 0 
	for i := 0 ; i < len(str) ; i++ {
		fmt.Println(str[i], len(str[i]))
		if len(str[i]) <= freeSpace {
			if freeSpace == len(str[i]) {
				freeSpace = freeSpace - (len(str[i])) 
			}
			freeSpace = freeSpace - (len(str[i])) + 1 
		}else {
			cost = cost + (freeSpace)
			freeSpace = k // newline
		}
	}
	fmt.Println(">>>  ", cost)
}
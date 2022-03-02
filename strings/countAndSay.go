package strings

import (
	"strconv"
	"fmt"
)

func CountAndSay(n int) {
	s := ""
	if n == 1 {
		s = s + "1"
	}

	if n == 2 {
		s = s +"1"+"1"
	}

	for i := 3 ; i <= n ; i++{
		temp := ""
		count := 1 
		for j := 1 ; j < len(s) ; j++ { //why you need an extra letter to added.
			if s[j-1] == s[j] {
				count ++
			}else {
				temp = temp + strconv.Itoa(count)
				s = s+temp
				count = 1 
			}
		} 
	}
	fmt.Println(s)
}
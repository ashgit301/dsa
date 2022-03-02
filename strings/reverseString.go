package strings

import (
	"fmt"
)

func ReverseString(s string)  { 
	end := len(s)-1
	op := ""
	for end >= 0 {
		op = op + string(s[end])
		end --
	}
	fmt.Println(op)
}
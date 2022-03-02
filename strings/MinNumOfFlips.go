package strings

import (
	"fmt"
)

func MinimumNumberOfFlips(s string)  {
	fmt.Println(MinX(ExpectedFlip(s,"0"), ExpectedFlip(s,"1")))
}



func ExpectedFlip(s string, char string) int {
	expected := char
	flipCount := 0 
	for i := 0 ; i < len(s) ; i ++ {
		if string(s[i]) != expected {
			flipCount++
		}
		if expected == "0" {
			expected = "1"
		}else {
			expected = "0"
		}
	}
	return flipCount
}

func MinX(x, y int) int {
    if x > y {
        return y
    }
    return x
}
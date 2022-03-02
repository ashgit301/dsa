package strings

import (
	"fmt"
)

func MinCharPalindrome(s string) {
	count := 0 
	for len(s) > 0 {
		if !IsPalindrom(0,len(s),s) {
			count++
			MinCharPalindrome(s[:len(s)-1])
		}else {
			break
		}
	}
	fmt.Println(count)
}


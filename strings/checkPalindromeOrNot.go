package strings

import (
	"fmt"
	"os"
)

func CheckPalindromeOrNot(s string) {
	start := 0 
	end := len(s)-1
	for start < end {
		if string(s[start]) != string(s[end]){
			os.Exit(2)
		}
		start++
		end--
	}
	fmt.Println("it is palindrome")
}
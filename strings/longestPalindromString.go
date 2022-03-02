package strings

import (
	"fmt"
)

func LongestPalindrom(s string) {
	maxLen := 0 
	var result string
	for i := 0 ; i < len(s) ; i ++ {
		for j := i ; j < len(s) ; j ++{
			if IsPalindrom(i,j,s) {
				fmt.Println(s[i:j+1]) //only j+2 will cause issue - panic
				if len(s[i:j+1]) > maxLen{
					maxLen = len(s[i:j+1])
					result = s[i:j+1]
				}
			}
		}
	}
	fmt.Println(result)
}

func IsPalindrom(i,j int, s string) bool {
	lo := i
	hi := j
	for lo < hi {
		if s[lo] == s[hi] {
			lo++
			hi--
		}else {
			return false
		}
	}
	return true
}

func LongestPalindromUsngRecursion(s string) {
	result := ""
	maxLen := 0
	palindromeHelper(s, "", &maxLen, &result, 0)
	fmt.Println(maxLen)
	fmt.Println(result)
}

func palindromeHelper(s string, temp string, maxLen *int, result *string, index int) {
	if index == len(s) {
		if IsPalindromRecur(temp) {
			if len(temp) > *maxLen {
				*maxLen = len(temp)
				*result = temp
				return
			}
		}
		return
	}
	if len(temp) > 0 && IsPalindromRecur(temp) {
		if len(temp) > *maxLen {
			*maxLen = len(temp)
			*result = temp
		}
	}
	palindromeHelper(s, temp+string(s[index]), maxLen, result, index+1)
	palindromeHelper(s, temp, maxLen, result, index+1)
	
}

func IsPalindromRecur(s string) bool{
	lo := 0
	hi := len(s)-1
	for lo < hi {
		if s[lo] == s[hi] {
			lo++
			hi--
		}else {
			return false
		}
	}
	return true
}
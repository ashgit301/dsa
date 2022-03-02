package strings

import (
	"fmt"
)

//word-by-word matching 

func LongestCommonPrefix(s []string) {
	ans := ""
	prefix := s[0]
	for i := 1 ; i < len(s) ; i ++ {
		prefix = FindPrefix(prefix,s[i])
		ans = prefix

	}
	fmt.Println(ans)
}

func FindPrefix(a,b string)string{
	i := 0 
	j := 0 
	result := ""
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = result+string(a[i])
		}
		i++
		j++
	}
	return result
}


//letter-by-letter 

func LongestCommonPrefix2(s []string) {
	result := ""
	minLen := len(s[0])
	for i := 1 ; i < len(s) ; i++ {
		if len(s[i]) < minLen {
			minLen = len(s[i])
		}
	}
	for i := 0 ; i < len(s) ; i ++ {
		curr := s[0][i]
		for j := 1 ; j < len(s) ; j ++ {
			if curr != s[j][0] {
				break
			}
		}
		result = result + string(curr)
	}
}
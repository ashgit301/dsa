package strings

import (
	"fmt"
)
//
func MinSwap(s string) int {
	stack := []string{}
	fmt.Println(stack)

	if len(s) % 2 != 0 {
		return -1
	}

	for i := 0 ; i < len(s) ; i ++ {
		if s[i] == '}' {
			if len(stack) != 0 && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			}else {
				stack = append(stack, string(s[i]))
			}
		}else {
			stack = append(stack, string(s[i]))
		}
	}
	fmt.Println(stack)

	reducedLen := len(stack)
	opening := 0 
	for len(stack) > 0 {
		if stack[len(stack)-1] == "}" {
			opening ++ 
			stack = stack[:len(stack)-1]
		}
	}
	// return ceil(m/2) + ceil(n/2) which is
    // actually equal to (m+n)/2 + n%2 when
    // m+n is even.
	return (reducedLen/2 +opening%2)
}
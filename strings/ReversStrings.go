package strings

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
    strarr := strings.Fields(s)
    strarr = reverse(strarr)
    return strings.Join(strarr, " ")
    
}

func reverse(a []string) []string {
    fmt.Println(">> ",a)
    lo := 0 
    hi := len(a) - 1
    for (lo < hi){
        a[lo],a[hi] = a[hi], a[lo]
        lo++
        hi--
    }
    return a
}
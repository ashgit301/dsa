package strings

import (
	"fmt"
)
//can use naive pattern matching 

func ZAlgorithm(s string, pattern string) {
	comp := pattern+"$"+s
	z := make([]int, len(comp))
	z[0]=0
	for i := 1 ; i < len(comp) ; i ++ {
		if comp[0] == comp[i] {
			count := 0
			j := 0 
			k := i 
			for j < len(comp) && k < len(comp) && comp[j] ==comp[k] {
				j++
				k++
				count++
			}
			z[i] = count
		}
	}
	fmt.Println(z)
}
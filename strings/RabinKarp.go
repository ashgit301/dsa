package strings

import (
	"fmt"
)

func NaivePatternMatching(Text, Pattern string) {
	N := len(Text)
	M := len(Pattern)
	for i := 0 ; i < N-M ; i ++ { //N_M cus max we need M letters
		j := 0
		for j = 0 ; j < M ; j ++ {
			if Text[i+j] != Pattern[j] {
				break
			}
		}
		if (j == M) {
			fmt.Println("Pattern found : ", i,j)
		}
		//pattern matches
	}
}


func RabinKarp(s string, pattern string){
	M := len(s)
	N := len(pattern)
	p := 0
	t := 0 
	//hash of pattern
	for i := 0 ; i < N ; i ++ {
		p = (p -'A')+1
	}
	//hash of first window
	for i := 0 ; i < N ; i ++ {
		t = (t -'A')+1
	}

	for i := 0 ; i < M-N ; i ++ {
		j := 0 
		if p == t { //num can match but letter might not 
			for j = 0 ; j < N ; j++{
				if s[i+j] != pattern[j] {
					break
				}
			}
			if (j == M) {
				fmt.Println("Pattern found : ", i,j)
			}
		}
		//next pattern
		t = (t - int(s[i]))+int(s[i+N]) //wrong

	}
}
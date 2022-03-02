package strings

import (
	"fmt"
)


func WordBreak(s string, wordDict []string) {
	ans := WordBreakHelper(s,wordDict)
	fmt.Println(ans)
}

func WordBreakHelper(s string, wordDict []string) bool {
	fmt.Println(s)
	if ArrContains(wordDict,s){
		return true
	}
	for i := 1 ; i < len(s) ; i ++ {
		currWord := s[0:i]
		if ArrContains(wordDict,currWord) && WordBreakHelper(s[i:], wordDict){ //easy. think about it 
			
			return true
		}
	}
	return false
}


func ArrContains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
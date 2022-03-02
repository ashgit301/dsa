package strings
import (
	"fmt"
)
func EditDistance(s1, s2 string) { //need to check how to memoize this solution. 
	ans := EditDistanceHelper(s1,s2,len(s1)-1,len(s2)-1)
	fmt.Println(ans)
}

func EditDistanceHelper(s1, s2 string, i,j int) int{

	if i == 0 {
		return j
	} else if j == 0 {
		return i
	}

	if s1[i] == s2[j] {
		return EditDistanceHelper(s1,s2,i-1,j-1)
	}

	insert := EditDistanceHelper(s1,s2,i,j-1) +1 
	delete := EditDistanceHelper(s1,s2,i-1,j) +1 
	replace := EditDistanceHelper(s1,s2,i-1,j-1) +1 
	return Min(insert,delete,replace)
}

func Min(a,b,c int)int {
	if a <=b && a <= c {
		return a 
	}else if b <=a && b <= c{
		return b
	}else {
		return c
	}
}


//easy with dynamic programming 



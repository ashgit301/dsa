package greedy

import (
	"fmt"
)

func SurviveInIslandOrNot(S,N,M int) {
	if ((N*6) < (M*7) && S > 6 ) || M > N { // cant survive on week (if several week) && cant survice on day 
		fmt.Println("NO")
	}
	days := (S*M)/N //take ceil value
	if (((M * S) % N) != 0){
		days++; //dont know why
	}
}
package strings

import (
	"sort"
	"strings"
)

func ValidShuffle(shuffle , a , b string) {
	temp := a+b
	if len(temp) != len(shuffle) {
		//false
	}
	sa := strings.Fields(shuffle)
	sb := strings.Fields(temp)
	sort.Strings(sa)
	sort.Strings(sb)
	//loop and check if "=="

}
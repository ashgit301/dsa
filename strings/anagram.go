package strings
import (
	"sort"
	"strings"
)
func Anagram(s1,s2 string) {
	sa := strings.Fields(s1)
	sb := strings.Fields(s2)
	sort.Strings(sa)
	sort.Strings(sb)
	i := 0 
	j := 0 
	for i < len(sa) && j < len(sb) {
		if sa[i] != sb[j] {
			break
		}
	}
	
}
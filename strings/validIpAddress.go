package strings

import (
	"fmt"
	"strconv"
)

func IPAddress(s string ) {
	path := make([]string, 4)
	ans := []string{}
	IPAddressHelper(path, 0, 0, s, &ans)
	fmt.Println(ans)
}

func IPAddressHelper(path []string, segment int, index int, s string, result *[]string ) {

	if segment == 3 && index == len(s) {
		fmt.Println("hello")
		*result = append(*result, path[0]+"."+path[1]+"."+path[2]+"."+path[3])
		return
	}

	if segment == 4 || index == len(s) {
		return
	} 

	for i := 1 ; i <= 3 && (i+index) < len(s) ; i ++ {
		element := s[index:i+index]
		intElement,_ := strconv.Atoi(element)
		fmt.Println(element)
		if intElement < 256 {
			path[segment] = element
			IPAddressHelper(path, segment+1,index+i,s,result)
			path[segment] = ""
		}
	}
}
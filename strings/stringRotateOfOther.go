package strings

import (
	"fmt"
	"strings"
)
func RotationOfOther(a,b string) {
	temp := a+b
	if strings.Contains(temp, b) {
		fmt.Println("Yes")
	}else {
		//exit
	}

}
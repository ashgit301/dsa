package linkedList

import (
	"fmt"
)

func Display(Head *Node) {
	for Head != nil {
		fmt.Println(Head.Val)
		Head = Head.Next
	}
}
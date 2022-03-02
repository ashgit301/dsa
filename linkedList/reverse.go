package linkedList

func Reverse(Head *Node) *Node {
	var prev *Node
	curr := Head
	nxt := Head
	for curr != nil {
		nxt = curr.Next
		curr.Next=prev
		prev = curr
		curr = nxt
	}
	return prev
}



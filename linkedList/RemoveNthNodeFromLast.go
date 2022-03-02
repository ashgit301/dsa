package linkedList

func RemoveNthNode(Head*Node, N int) *Node {
	count := 1
	curr := Head
	for curr != nil {
		curr = curr.Next
		count++
	}
	target := count - N
	curr = Head
	for target == 1 {
		curr = curr.Next
		target --
	}
	curr.Next = curr.Next.Next
	return Head
}

func RemoveNthNode2(Head*Node, N int) *Node {
	var start *Node
	start.Next = Head
	slow := start
	fast := start
	for i := 1 ; i <= N ; i ++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return start.Next
	
}

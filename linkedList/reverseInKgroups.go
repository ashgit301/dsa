package linkedList

func ReverseInKGroups(head *Node, k int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy *Node
	dummy.Next = head
	curr := dummy
	next := dummy
	prev := dummy
	count := 0 
	for curr.Next != nil {
		count = count + 1
		curr = curr.Next
	}
	for count >= k {
		for i := 1 ; i < k ; i++ {
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
			next = curr.Next

		} 
		count = count - k 
		prev = curr
	}
	return dummy.Next
}
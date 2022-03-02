package linkedList

func RemovedDuplicatesFromLL(head *Node)*Node {
	curr := head
	for curr != nil && curr.Next != nil {
		for curr.Val == curr.Next.Val && curr.Next != nil{
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
	return head
}
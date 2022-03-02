package linkedList

func RotateLinkedList(head *Node, k int) *Node {
	curr := head
	length := 1
	for curr.Next != nil {
		curr = curr.Next
	}
	k = length % k 
	k = length - k 
	curr.Next = head
	for k != 0 {
		curr = curr.Next
		k--
	}
	head = curr.Next
	curr.Next = nil
	return head
}
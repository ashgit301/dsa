package linkedList

func PalindromeOrNot(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = Reverse(slow.Next)
	slow = slow.Next
	for slow.Next != nil {
		if slow.Val != head.Val{
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}
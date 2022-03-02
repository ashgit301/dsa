package linkedList

func DetectCycle(Head *Node) bool {
	if Head == nil || Head.Next == nil {
		return false
	}
	slow := Head
	fast := Head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
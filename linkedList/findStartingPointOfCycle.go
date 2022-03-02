package linkedList

func CycleStartingPoint(head *Node) *Node {
	if (head == nil || head.Next == nil){
		return head
	}
	slow := head
	fast := head
	entry := head
	for fast != nil || fast.Next != nil {
		slow =slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = slow.Next
			entry = entry.Next
			if slow == entry {
				return entry
			}
		}
	}
	return nil
}
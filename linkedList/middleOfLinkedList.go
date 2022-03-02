package linkedList

func MiddleElement(Head *Node) *Node {
	slow := Head
	fast := Head
	for fast != nil || fast.Next != nil {
		slow =slow.Next
		fast = fast.Next.Next
	}
	return slow
}
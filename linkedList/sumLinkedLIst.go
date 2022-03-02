package linkedList

func SumTwoLinkedList(l1,l2 *Node) *Node {
	carry := 0
	var dummy *Node
	tmp := dummy
	for (l1 != nil || l2 != nil) {
		sum := 0
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		sum = sum + carry
		carry = sum / 10
		newNode := &Node{}
		newNode.Val = sum%10
		tmp.Next = newNode
		tmp = tmp.Next
	}
	return dummy.Next
}
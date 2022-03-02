package linkedList

func RemoveDuplicatedUnsorted(head *Node) *Node {
	curr := head
	for curr != nil && curr.Next != nil {
		tmp := curr
		for tmp.Next != nil {
			if tmp.Val == tmp.Next.Val {
				tmp.Next = tmp.Next.Next
			}else {
				tmp = tmp.Next
			}
		}
		curr = curr.Next
	}
	return head
}

func RemoveDuplicatedUnsorted2(head *Node) *Node {
	var mpp map[*Node]bool
	curr :=head
	var prev *Node
	prev = nil
	for curr != nil && curr.Next != nil {
		if _,ok := mpp[curr]; ok {
			prev.Next = curr.Next
		}else {
			mpp[curr] = true
			prev = curr
		}
		curr = prev.Next
	}
	return head
}
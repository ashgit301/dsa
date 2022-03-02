package linkedList

func MergeLinkedList(l1, l2 *Node) *Node {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *Node
	dummy := head

	for l1 != nil || l2 != nil {
		if l1.Val < l2.Val {
			dummy.Next = l1
			l1= l1.Next
		}else if l2.Val < l1.Val {
			dummy.Next = l2
			l2= l2.Next
		}
		dummy = dummy.Next
	}
	for l1 != nil {
		dummy.Next = l1
		l1 = l1.Next
		dummy = dummy.Next
	}
	for l2 != nil {
		dummy.Next = l2
		l2 = l2.Next
		dummy = dummy.Next
	}
	return head
}

func MergeLinkedList2(l1,l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		tmp := l1
		l1 = l2 
		l2 = tmp
	}
	res := l1
	for l1 != nil && l2 != nil {
		tmp := l1
		for l1 != nil && l1.Val <= l2.Val {
			tmp = l1
			l1 = l1.Next
		}
		tmp.Next = l2
		swap := l1
		l1 = l2 
		l2 = swap
	}
	return res
}
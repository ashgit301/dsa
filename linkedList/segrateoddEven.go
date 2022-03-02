package linkedList

func SegregateOddEven(head *Node)*Node {
	var (
		evenStart *Node
		evenEnd *Node
		oddStart *Node
		oddEnd *Node
	)
	evenEnd = nil
	oddStart = nil
	evenStart = nil
	oddEnd = nil
	curr := head
	for curr.Next != nil {
		if curr.Val % 2 == 0 {
			if evenStart == nil {
				evenStart = curr
				evenEnd = evenStart
			}else {
				evenEnd.Next = curr
				evenEnd = evenEnd.Next
			}
		}else {
			if oddStart == nil {
				oddStart = curr
				oddEnd = oddEnd
			}else{
				oddEnd.Next = curr
				oddEnd = oddEnd.Next
			}
		}
	curr = curr.Next
	}
	oddEnd.Next = evenStart
	evenEnd = nil
	head = oddStart
	return head
}
package linkedList

func Segregate(head *Node)*Node {
	count := []int{0,0,0}
	curr := head
	for curr.Next != nil {
		count[curr.Val]++
	}
	curr = head
	i := 0 
	for curr != nil {
		if count[i] == 0 {
			i++
		}else {
			curr.Val = count[i]
			count[i]--
			curr = curr.Next
		}
	}
	return head
}
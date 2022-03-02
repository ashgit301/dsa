package linkedList

func Intersection(l1,l2 *Node) int {
	mpp := map[*Node]int{}
	for l1 != nil {
		mpp[l1] = l1.Val
		l1 = l1.Next
	}
	for l2 != nil {
		if _,ok := mpp[l2]; ok {
			return l2.Val
		}
	}
	return -1
}


//calculate the len of the longest linked list 
//calculate the difference between the lens
//move the longest linked list to diff
//move both l1 & l2 by one point
//if both l1 & l2 eqaul, return value
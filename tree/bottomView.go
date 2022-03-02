package tree

type BottomHD struct{
	Node *Node
	hd int
}
func BottomView(root *Node, result *[]int){
	mpp := map[int]*Node{}
	q := []BottomHD{}
	initial := BottomHD{
		Node : root,
		hd 	 : 0,
	}
	q = append(q,initial)
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i ++ {
			curr := q[0]
			q = q[1:]
			currHD := curr.hd
			mpp[currHD] = curr.Node
			if curr.Node.Left != nil {
				initial := BottomHD{
					Node : curr.Node.Left,
					hd 	 : currHD-1,
				}	
				q = append(q,initial)						
			}
			if curr.Node.Right != nil {
				initial := BottomHD{
					Node : curr.Node.Right,
					hd 	 : currHD+1,
				}	
				q = append(q,initial)						
			}
		}
	}
	for _,Val := range mpp {
		*result = append(*result,Val.Val)
	}
}
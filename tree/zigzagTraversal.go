package tree

func ZigZag(root *Node, result *[]int) {
	zigzag := false
	currStack := []*Node{root}
	nextStack := []*Node{}
	_ = nextStack
	for len(currStack) != 0 {
		length := len(currStack)
		for i := length-1 ; i >= 0 ; i++ {
			*result = append(*result, currStack[i].Val)
			curr := currStack[0]
			currStack = currStack[1:]
			if !zigzag {
				if curr != nil {
					if curr.Left != nil {
						nextStack = append(currStack, curr.Left)
					}
					if curr.Right != nil {
						nextStack = append(currStack, curr.Right)
					}
				}
			} else {
				if curr != nil {
					if curr.Right != nil {
						nextStack = append(currStack, curr.Right)
					}
					if curr.Left != nil {
						nextStack = append(currStack, curr.Left)
					}
				}
			}
		}
		//swap currStack with nextStack
	}

}
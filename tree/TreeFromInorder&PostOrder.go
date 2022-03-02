package tree

func PostOrder_Inorder(Postorder, Inroder []int)*Node {
	curr := Postorder[0]
	tmp := 0 
	for i := 0 ; i < len(Inroder) ; i++ {
		if curr == Inroder[i] {
			tmp = i
		}
	}
	return  &Node{
		Val:   curr,
		Left:  PostOrder_Inorder(Postorder[:tmp], Inroder[:tmp]),
		Right: PostOrder_Inorder(Postorder[tmp:len(Postorder)-1], Inroder[tmp+1:]),
	}
}


//In - 4251637
//post - 4526731




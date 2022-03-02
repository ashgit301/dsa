package tree


func RightView(root *Node,result *[]int) {
    q := []*Node{root}
    for len(q) > 0 {
        length := len(q)
        for i := 0 ; i < length ; i++ {
            curr := q[0]
            q = q[1:]
            if i == length-1 {
                *result = append(*result,curr.Val)
            }
            if curr.Left != nil {
                q = append(q,curr.Left)
            }
            if curr.Right != nil {
                q = append(q,curr.Right)
            }
        }
    }
}
   
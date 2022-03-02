package tree

func isSameTree(p *Node, q *Node) bool {
    
    if (p == nil && q == nil) {
        return true
    } else if (p == nil || q == nil) || (p.Val != q.Val) {
        return false
    }
    
    return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right) 
}
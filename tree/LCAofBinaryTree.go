
package tree


func lowestCommonAncestor(root, p, q *Node) *Node {

    if root == nil || root.Val == p.Val || root.Val == q.Val {   //it actually returns nil also as root <trace it>
        return root
    }
    
    leftLca := lowestCommonAncestor(root.Left, p, q)
    rightLca := lowestCommonAncestor(root.Right, p, q)
    
    if leftLca != nil {
        if rightLca != nil {
            return root
        }
        return leftLca
    }
    return rightLca
}
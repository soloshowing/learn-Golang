package tree

func (node *Node) Traverse() {
	node.TraverseFunc(func(nd *Node) {
		nd.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

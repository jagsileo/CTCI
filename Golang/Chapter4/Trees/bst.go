package bst

type Node struct {
	Item  int
	Left  *Node
	Right *Node
}

func bstOperations(arr []int) *Node {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	bstNode := buildBST(arr)
	bstNode.printNode()
	return bstNode
}

func buildBST(arr []int) *Node {
	bst := Node{}
	bstCopy := &bst
	for _, num := range arr {
		bstCopy.insertNode(num)
	}
	return &bst
}

func (node *Node) insertNode(num int) {
	if node == nil {
		node = &Node{Item: num, Left: nil, Right: nil}
		return
	}

	if node.Item == 0 {
		node.Item = num
		return
	}

	if num < node.Item {
		(node.Left).insertNode(num)
	} else if num > node.Item {
		(node.Right).insertNode(num)
	}

}

func (node *Node) printNode() {
	if node == nil {
		return
	}
	print(node.Item)
	print("\n")
	(node.Left).printNode()
	(node.Right).printNode()

}

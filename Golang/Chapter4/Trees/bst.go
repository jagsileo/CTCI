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
		bstCopy = bstCopy.insertNode(num)
	}
	return bstCopy
}

func (node *Node) insertNode(num int) *Node {
	if node == nil {
		node = &Node{Item: num, Left: nil, Right: nil}
		return node
	}

	if node.Item == 0 {
		node.Item = num
		return node
	}

	if num < node.Item {
		node.Left = (node.Left).insertNode(num)
	} else if num > node.Item {
		node.Right = (node.Right).insertNode(num)
	}
	return node

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

func (node *Node) calcHeight() int {
	if node == nil {
		return 0
	}

	return 1 + max(node.Left.calcHeight(), node.Right.calcHeight())
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

// func avlTree(node *Node) *Node {
// 	if node == nil {
// 		return nil
// 	}

// 	hDiff := node.Right.calcHeight() - node.Left.calcHeight()

// }

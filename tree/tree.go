package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

var currentValue int

func BuildRandomTree(elements int) *Node {
	currentValue = 0
	return buildTree(elements)
}

func buildTree(elements int) *Node {
	if elements <= 0 {
		return nil
	}
	var lElements, rElements int
	elements--
	lElements = elements / 2
	if (elements % 2) == 0 {
		rElements = elements / 2
	} else {
		rElements = (elements / 2) + 1
	}
	currentValue++
	node := &Node{
		Value: currentValue,
	}
	node.Left = buildTree(lElements)
	node.Right = buildTree(rElements)

	return node
}

func PrintTreeNLR(root *Node) {
	if root != nil {
		fmt.Printf("%d", root.Value)
		PrintTreeNLR(root.Left)
		PrintTreeNLR(root.Right)
	}
}

func PrintTreeLNR(root *Node) {
	if root != nil {
		PrintTreeLNR(root.Left)
		fmt.Printf("%d", root.Value)
		PrintTreeLNR(root.Right)
	}
}

func PrintTreeRNL(root *Node) {
	if root != nil {
		PrintTreeRNL(root.Right)
		fmt.Printf("%d", root.Value)
		PrintTreeRNL(root.Left)
	}
}

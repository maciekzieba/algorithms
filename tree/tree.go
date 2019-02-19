package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type stack struct {
	items []*Node
}

func newNodeStack() *stack {
	return &stack{items: make([]*Node, 0)}
}

func (s *stack) push(node *Node) {
	s.items = append(s.items, node)
}

func (s *stack) pop() *Node {
	if len(s.items) == 0 {
		return nil
	}
	node := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return node
}

func (s stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
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

func PrinttIterative(root *Node) {
	if root == nil {
		return
	}
	nodesStack := newNodeStack()
	nodesStack.push(root)

	for nodesStack.isEmpty() == false {
		node := nodesStack.pop()
		fmt.Printf("%d ", node.Value)
		if node.Right != nil {
			nodesStack.push(node.Right)
		}
		if node.Left != nil {
			nodesStack.push(node.Left)
		}
	}

}

func PrintTreeNLR(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Value)
		PrintTreeNLR(root.Left)
		PrintTreeNLR(root.Right)
	}
}

func PrintTreeLNR(root *Node) {
	if root != nil {
		PrintTreeLNR(root.Left)
		fmt.Printf("%d ", root.Value)
		PrintTreeLNR(root.Right)
	}
}

func PrintTreeRNL(root *Node) {
	if root != nil {
		PrintTreeRNL(root.Right)
		fmt.Printf("%d ", root.Value)
		PrintTreeRNL(root.Left)
	}
}

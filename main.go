package main

import (
	"fmt"
	"github.com/maciekzieba/algorithms/tree"
)

func main() {
	root := tree.BuildRandomTree(18)
	fmt.Println("NLR Tree:")
	tree.PrintTreeNLR(root)
	fmt.Println()
	fmt.Println("LNR Tree:")
	tree.PrintTreeLNR(root)
	fmt.Println()
	fmt.Println("RNL Tree:")
	tree.PrintTreeRNL(root)
	fmt.Println()
	fmt.Println("iterative Tree:")
	tree.PrinttIterative(root)
	fmt.Println()
}

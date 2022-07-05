package main

import (
	"fmt"
)

type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

func (root *TreeNode) PreOrderTraversal() {

	if root == nil {
		return
	}
	fmt.Printf(root.data)
	root.left.PreOrderTraversal()
	root.right.PreOrderTraversal()
}

func (root *TreeNode) PostOrderTraversal() {

	if root == nil {
		return
	}

	root.left.PreOrderTraversal()
	root.right.PreOrderTraversal()
	fmt.Printf(root.data)
}
func makeNode(data string) *TreeNode {
	return &TreeNode{data, nil, nil}
}
func main() {

	//                    +
	//                /      \
	//              a         -
	//                      /   \
	//                     *     c
	//                   /  \
	//                  *    ^
	//                / \   / \
	//               d  e  f  g
	//         a + ((d*e * f^g)-c)

	root := makeNode("+")
	root.left = makeNode("a")
	root.right = makeNode("-")
	root.right.left = makeNode("*")
	root.right.right = makeNode("c")
	root.right.left.left = makeNode("*")
	root.right.left.right = makeNode("^")
	root.right.left.left.left = makeNode("d")
	root.right.left.left.right = makeNode("e")
	root.right.left.right.left = makeNode("f")
	root.right.left.right.right = makeNode("g")

	fmt.Println("PreOrder Traversal")
	root.PreOrderTraversal()
	fmt.Println("")
	fmt.Println("PostOrder Traversal")
	root.PostOrderTraversal()

}

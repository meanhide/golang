package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)

	RemoveNode(root)
	PrintNodes(root)
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) {
	if node == root {
		root = root.next
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev.prev.next
	}
	if node == tail {
		prev.next
	}
	prev.next = prev.next.next

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)

}

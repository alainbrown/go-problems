package main

import (
	"fmt"
)

func main() {
	sp := &Node{}
	p := &Node{&Node{&Node{sp}}}
	n := &Node{&Node{p}}
	m := &Node{&Node{&Node{p}}}
	fmt.Println(p == lac(n, m))
}

type Node struct {
	Parent *Node
}

func lac(first, second *Node) *Node {
	fD := depth(first)
	sD := depth(second)
	depth := min(fD, sD)
	first = swim(first, depth, fD)
	second = swim(second, depth, sD)
	for first != second {
		first = first.Parent
		second = second.Parent
	}
	return first
}

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}

func swim(node *Node, depth int, current int) *Node {
	if current <= depth {
		return node
	}
	return swim(node.Parent, depth, current-1)
}

func depth(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + depth(node.Parent)
}

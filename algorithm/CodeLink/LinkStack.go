package main

type Node struct {
	data  interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.pNext == nil
}

func (n *Node) Push(data interface{}) {
	newnode := &Node{data: data}
	newnode.pNext = n.pNext
	n.pNext = newnode
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() {
		return nil
	}
	value := n.pNext.data
	n.pNext = n.pNext.pNext
	return value
}

func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext
		length++
	}
	return length
}

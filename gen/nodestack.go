package gen

import (
	"go/ast"
)

type NodeStack interface {
	Push(*ast.Node) NodeStack
	Pop() (*ast.Node, *NodeStack)
	Peek() *ast.Node
	PeekDeep(depth int) *ast.Node
	Size() int
}

type nodeStack struct {
	data []ast.Node
}

func NewNodeStack() NodeStack {
	return nodeStack{}
}

func (ns nodeStack) Push(n *ast.Node) NodeStack {
	nsRes := nodeStack{}
	if ns.data == nil {
		nsRes.data = []ast.Node{*n}
	} else {
		nsRes.data = append(ns.data, *n)
	}
	return &nsRes
}

func (ns nodeStack) Pop() (*ast.Node, *NodeStack) {
	if len(ns.data) > 0 {
		size := len(ns.data)
		var nsRes interface{} = nodeStack{data: ns.data[0:size-1]}
		nsResInt := nsRes.(NodeStack)
		return &(ns.data[size -1]), &nsResInt
	} else {
		return nil, nil
	}
}

func (ns nodeStack) peek(depth int) *ast.Node {
	if depth < len(ns.data) {
		return &(ns.data[len(ns.data)-depth-1])
	} else {
		return nil
	}
}

func (ns nodeStack) Peek() *ast.Node {
	return ns.peek(0)
}

func (ns nodeStack) PeekDeep(depth int) *ast.Node {
	return ns.peek(depth)
}

func (ns nodeStack) Size() int {
	return len(ns.data)
}

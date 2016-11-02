package gen

import (
	"go/ast"
	"container/list"
)

type nodeHandler struct {
	stack *NodeStack
	callback func(stack NodeStack, params map[string]string)
}

type astWalker struct {
	stack *NodeStack
	xformTargets list.List
}

func (state *astWalker) walk (n ast.Node) bool {
	x := (*state.stack).Push(&n)
	state.stack = &x
	nh, res := checkNode(n, state.stack)
	if nh != nil {
		state.xformTargets.PushBack(nh)
	}
	_, state.stack = (*state.stack).Pop()
	return res;
}

func checkNode(n ast.Node, stack *NodeStack) (*nodeHandler, bool) {
	nh := nodeHandler{stack: stack}
	switch x := n.(type) {
	case *ast.SelectorExpr:
		switch y := x.X.(type) {
		case *ast.Ident:
			if y.Name == "_t_" {
				nh.callback = GenericGenerator
				return &nh, true
			}
		}
	case *ast.CallExpr:
		switch y := x.Fun.(type) {
		case *ast.SelectorExpr:
			if y.Sel.Name == "check" {
				return &nh, true
			}
		}
	}
	return nil, true
}

func GrabNodes (n ast.Node) list.List {
	stack := NewNodeStack()

	walker := astWalker{stack: &stack}

	ast.Inspect(n, walker.walk)
	return walker.xformTargets
}


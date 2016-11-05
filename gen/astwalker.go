package gen

import (
	"go/ast"
	"container/list"
	"flag"
)

type Matcher func(stack NodeStack, options flag.FlagSet) bool

type NodeHandler struct {
	stack *NodeStack
	manipulate Manipulator
	getService func (name string) interface{}
}

type NodeAnalyzer struct {
	match Matcher
	manipulate Manipulator
	getService func (name string) interface{}
}

type astWalker struct {
	options flag.FlagSet
	stack *NodeStack
	xformTargets list.List
}

var generators = []NodeAnalyzer {
	NodeAnalyzer{match: MatchCheckCall, manipulate: ReplaceCheckCall},
	NodeAnalyzer{match: MatchGenericTypeRef, manipulate: GenericGenerator},
}

func (state *astWalker) walk (n ast.Node) bool {
	x := (*state.stack).Push(&n)
	state.stack = &x
	nh, res := checkNode(n, state.stack, state.options)
	if nh != nil {
		state.xformTargets.PushBack(nh)
	}
	_, state.stack = (*state.stack).Pop()
	return res;
}

func checkNode(n ast.Node, stack *NodeStack, options flag.FlagSet) (*NodeHandler, bool) {
	for _,generator := range generators {
		if generator.match(*stack, options) {
			nh := NodeHandler{
				stack: stack,
				manipulate: generator.manipulate,
				getService: generator.getService,
			}
			return &nh, true
		}
	}
	//nh := nodeHandler{stack: stack}
	//switch x := n.(type) {
	//case *ast.SelectorExpr:
	//	switch y := x.X.(type) {
	//	case *ast.Ident:
	//		if y.Name == "_t_" {
	//			nh.callback = GenericGenerator
	//			return &nh, true
	//		}
	//	}
	//case *ast.CallExpr:
	//	switch y := x.Fun.(type) {
	//	case *ast.SelectorExpr:
	//		if y.Sel.Name == "check" {
	//			nh.callback = CheckGenerator
	//			return &nh, true
	//		}
	//	}
	//}
	return nil, true
}

func toNhArray(list list.List) []NodeHandler {
	res := make([]NodeHandler, list.Len())
	n := 0
	for e := list.Front(); e != nil; e = e.Next() {
		res[n] = e.Value.(NodeHandler)
		n = n + 1
	}
	return res;
}

func GrabNodes (n ast.Node, options flag.FlagSet) TransformationPlan {
	stack := NewNodeStack()

	walker := astWalker{stack: &stack, options: options}

	ast.Inspect(n, walker.walk)
	targets := toNhArray(walker.xformTargets)
	return CreateTransformationPlan(targets, options)
}


package gen

import (
	"go/ast"
	"flag"
)

func MatchGenericTypeRef(stack NodeStack, options flag.FlagSet) bool {
	var top interface{} = stack.Peek()
	switch x := top.(type) {
	case *ast.SelectorExpr:
		switch y := x.X.(type) {
		case *ast.Ident:
			if y.Name == "_t_" {
				return true
			}
		}
	}
	return false;
}

func GenericGenerator(stack NodeStack, context TransformContext) error {
	var selector interface{} = stack.Peek()
	switch x := selector.(type) {
	case *ast.SelectorExpr:
		switch y := x.X.(type) {
		case *ast.Ident:
			y.Name = "foo"
		}
		x.Sel.Name = "bar"
	}

	return nil
}
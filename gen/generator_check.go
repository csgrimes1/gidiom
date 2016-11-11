package gen

import (
	"go/ast"
	"flag"
)

func MatchCheckCall(stack NodeStack, options flag.FlagSet) MatchResult {
	var top interface{} = stack.Peek()
	switch x := top.(type) {
	case *ast.CallExpr:
		switch y := x.Fun.(type) {
		case *ast.SelectorExpr:
			if y.Sel.Name == "check" {
				return Match
			}
		}
	}
	return NoMatch
}


func ReplaceCheckCall(stack NodeStack, context TransformContext) error {

	return nil
}

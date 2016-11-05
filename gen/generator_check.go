package gen

import (
	"go/ast"
	"flag"
	"fmt"
)

func MatchCheckCall(stack NodeStack, options flag.FlagSet) bool {
		fmt.Println(stack.Size())
	var top interface{} = stack.Peek()
	switch x := top.(type) {
	case *ast.CallExpr:
		switch y := x.Fun.(type) {
		case *ast.SelectorExpr:
			if y.Sel.Name == "check" {
				return true
			}
		}
	}
	return false
}


func ReplaceCheckCall(stack NodeStack, context TransformContext) error {

	return nil
}

package gen

import (
	"go/ast"
	"go/parser"
	//"go/printer"
	"go/token"
	"fmt"
	//"os"
	//"reflect"
	"container/list"
)

type VisitorBuilder struct {
	result, parent ast.Node;
	stack list.List
}

func (v VisitorBuilder) Visit(n ast.Node) ast.Visitor {
	v.result = n
	switch x := n.(type) {
	case *ast.CallExpr:
		//See if the method name is 'check'
		switch x2 := x.Fun.(type) {
		case *ast.SelectorExpr:
			fmt.Println(x2.Sel.Name)
		}
	}
	return VisitorBuilder{parent: n}
}

func (v VisitorBuilder) hola (n ast.Node) bool {

	v.stack.PushBack(n)
	return true;
}


func Compile(fileName string)  {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ast.Print(fset, f)
	// Inspect the AST and print all identifiers and literals.
	//ast.Inspect(f, func(n ast.Node) bool {
	//	var s string
	//	switch x := n.(type) {
	//	case *ast.BasicLit:
	//		s = x.Value
	//		//return false
	//	case *ast.Ident:
	//		s = x.Name
	//	}
	//	if s != "" {
	//		fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
	//	}
	//	return true
	//})
	//
	//fmt.Println("====================")
	//printer.Fprint(os.Stdout, fset, f)
	//
	v := VisitorBuilder{}
	//ast.Walk(v, f)

	ast.Print(fset, f)
	fmt.Println("====================")
	ast.Inspect(f, v.hola)
}

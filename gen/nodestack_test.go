package gen_test

import (
	"github.com/csgrimes1/gidiom/gen"
	"testing"
	"go/ast"
	"github.com/stretchr/testify/assert"
)

func TestMapping(t *testing.T) {
	stack1 := gen.NewNodeStack()

	var n1 ast.Node = new(ast.BadExpr)
	stack2 := stack1.Push(&n1)
	var n2 ast.Node = new(ast.BadStmt)
	stack3 := stack2.Push(&n2)

	assert.Equal(t, &n1, stack2.Peek())
	assert.Equal(t, &n2, stack3.Peek())
	nParent := stack3.PeekDeep(1)
	nilAncestor := stack3.PeekDeep(2)
	assert.Equal(t, &n1, nParent)
	assert.Nil(t, nilAncestor)

	n3, stack4 := stack3.Pop()
	assert.Equal(t, n3, &n2)
	assert.Equal(t, 1, (*stack4).Size())

	n4, stack5 := (*stack4).Pop()
	assert.Equal(t, n4, &n1)
	assert.Equal(t, 0, (*stack5).Size())

	n5, stack6 := (*stack5).Pop()
	assert.Nil(t, n5)
	assert.Nil(t, stack6)
}

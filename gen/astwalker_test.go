package gen_test

import (
	"github.com/csgrimes1/gidiom/gen"
	"testing"
	"go/parser"
	"go/token"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestGrabNodes(t *testing.T) {
	codefile := os.ExpandEnv("${GOPATH}/src/github.com/csgrimes1/gidiom/testdata/test.gid")
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, codefile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	targets := gen.GrabNodes(f)
	assert.Equal(t, 2, targets.Len())
}

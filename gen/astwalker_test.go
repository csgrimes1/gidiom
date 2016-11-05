package gen_test

import (
	"github.com/csgrimes1/gidiom/gen"
	"testing"
	"go/parser"
	"go/token"
	"github.com/stretchr/testify/assert"
	"os"
	"flag"
)

func TestGrabNodes(t *testing.T) {
	codefile := os.ExpandEnv("${GOPATH}/src/github.com/csgrimes1/gidiom/testdata/test.gid")
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, codefile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	options := flag.NewFlagSet("", flag.ContinueOnError)
	optcopy := *options
	targets := gen.GrabNodes(f, optcopy)
	assert.Equal(t, 2, targets.TransformationCount())// targets.Len())
}

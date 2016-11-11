package core_test

import (
	"github.com/csgrimes1/gidiom/core"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	any := core.CreateAny(1)
	assert.Equal(t, 1, any.RawValue())
}


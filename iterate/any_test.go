package iterate_test

import (
	"github.com/csgrimes1/gidiom/iterate"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	any := iterate.CreateAny(1)
	assert.Equal(t, 1, any.RawValue())
}


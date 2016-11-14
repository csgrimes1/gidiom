package iterate_test

import (
	"github.com/csgrimes1/gidiom/iterate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapIteration(t *testing.T) {
	var m = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	slice := iterate.CreateMapIterator(m).
		ToTypedSlice(iterate.MapPairType).([]iterate.MapPair)
	assert.Equal(t, 3, len(slice))
}

package iterate_test

import (
	"container/list"
	"github.com/csgrimes1/gidiom/iterate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardIteration(t *testing.T) {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	slice := iterate.CreateListIterator(list, true).ToSlice()
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, iterate.CreateAny(1), slice[0])
	assert.Equal(t, iterate.CreateAny(2), slice[1])
	assert.Equal(t, iterate.CreateAny(3), slice[2])
}

func TestReverseIteration(t *testing.T) {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	slice := iterate.CreateListIterator(list, false).ToSlice()
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, iterate.CreateAny(3), slice[0])
	assert.Equal(t, iterate.CreateAny(2), slice[1])
	assert.Equal(t, iterate.CreateAny(1), slice[2])
}

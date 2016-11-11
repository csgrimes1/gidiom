package core_test

import (
	"github.com/csgrimes1/gidiom/core"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCount0(t *testing.T) {
	it := core.CreateArrayIterator([]core.Any{})
	assert.Equal(t, 0, it.Count())
}

func TestCount1(t *testing.T) {
	it := core.CreateArrayIterator([]core.Any{core.CreateAny(1)})
	assert.Equal(t, 1, it.Count())
}

func TestCount3(t *testing.T) {
	it := core.CreateArrayIterator([]core.Any{core.CreateAny(1), core.CreateAny(2), core.CreateAny(3)})
	assert.Equal(t, 3, it.Count())
}

func TestMap(t *testing.T) {
	it := core.CreateArrayIterator([]core.Any{core.CreateAny(1), core.CreateAny(2), core.CreateAny(3)}).Map(func(num core.Any) core.Any {
		val := num.RawValue().(int)
		return core.CreateAny(val +1)
	})
	assert.Equal(t, 3, it.Count())
	slice := it.ToSlice()
	assert.Equal(t, 3, len(slice))

	expect := []core.Any{core.CreateAny(2), core.CreateAny(3), core.CreateAny(4)}
	assert.Equal(t, expect, slice)
}

package iterate_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/csgrimes1/gidiom/iterate"
	"reflect"
	"strconv"
	"testing"
)

func TestCount0(t *testing.T) {
	it := iterate.CreateArrayIterator([]iterate.Any{})
	assert.Equal(t, 0, it.Count())
}

func TestCount1(t *testing.T) {
	it := iterate.CreateArrayIterator([]iterate.Any{iterate.CreateAny(1)})
	assert.Equal(t, 1, it.Count())
}

func TestCount3(t *testing.T) {
	//it := iterate.CreateArrayIterator([]iterate.Any{iterate.CreateAny(1), iterate.CreateAny(2), iterate.CreateAny(3)})
	it := iterate.Start([]int{4, 5, 6})
	assert.Equal(t, 3, it.Count())
}

func TestMap(t *testing.T) {
	it := iterate.Start([]int{1, 2, 3}).Map(func(num iterate.Any) iterate.Any {
		val := num.RawValue().(int)
		return iterate.CreateAny(val+1)
	})
	assert.Equal(t, 3, it.Count())
	slice := it.ToSlice()
	assert.Equal(t, 3, len(slice))

	expect := []iterate.Any{iterate.CreateAny(2), iterate.CreateAny(3), iterate.CreateAny(4)}
	assert.Equal(t, expect, slice)
}

func TestMAP(t *testing.T) {
	it := iterate.Start([]int{10, 11, 12}).Map(iterate.MAP(func(n int) string {
		return strconv.Itoa(n)
	}))
	slice := it.ToSlice()
	assert.Equal(t, "10", slice[0].RawValue())
	assert.Equal(t, "11", slice[1].RawValue())
	assert.Equal(t, "12", slice[2].RawValue())
}

func TestMaybe(t *testing.T) {
	it := iterate.Start("test")
	assert.Equal(t, 1, it.Count())
}

func TestFilter(t *testing.T) {
	slice := iterate.Start([]int{1, 10, 3, 5, 200, -1}).Filter(iterate.PRED(func(n int) bool {
		return n < 10
	})).ToSlice()
	assert.Equal(t, 4, len(slice))
	assert.Equal(t, 1, slice[0].RawValue())
	assert.Equal(t, 3, slice[1].RawValue())
	assert.Equal(t, 5, slice[2].RawValue())
	assert.Equal(t, -1, slice[3].RawValue())
}

func fib (trail []int) ([]int, int) {
	next := trail[0] + trail[1]
	return []int{trail[1], next}, next
}

func fibonacci() iterate.Iterator {
	startingContext := iterate.CreateAny([]int{0, 1})
	return iterate.CreateSequence(startingContext, iterate.GENERATOR(fib))
}

func TestTake(t *testing.T) {
	slice := fibonacci().
		Take(5).
		ToSlice()
	assert.Equal(t, 5, len(slice))
	assert.Equal(t, 1, slice[0].RawValue())
	assert.Equal(t, 2, slice[1].RawValue())
	assert.Equal(t, 3, slice[2].RawValue())
	assert.Equal(t, 5, slice[3].RawValue())
	assert.Equal(t, 8, slice[4].RawValue())
}

func TestReduce(t *testing.T) {
	sum := fibonacci().
		Take(5).
		Reduce(iterate.CreateAny(0), iterate.REDUCER(func(accum int, element int) int {
			return accum + element
		}))

	assert.Equal(t, 19, sum.RawValue())
}

func TestTypedSliceM(t *testing.T) {
	converter := func(n int) string {
		return strconv.Itoa(n)
	}
	slice := fibonacci().
		Take(5).
		ToTypedSliceM(reflect.TypeOf(""), iterate.MAP(converter))

	assert.Equal(t, []string{"1","2","3","5","8"}, slice)
}

func TestTypedSlice(t *testing.T) {
	slice := fibonacci().
		Take(5).
		ToTypedSlice(reflect.TypeOf(0))

	assert.Equal(t, []int{1, 2, 3, 5, 8}, slice)
}

func TestIdentityMapping(t *testing.T) {
	res := iterate.IdentityMapping(iterate.CreateAny(200))
	assert.Equal(t, 200, res.RawValue())
}
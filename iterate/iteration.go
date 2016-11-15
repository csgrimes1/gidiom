package iterate

import (
	"container/list"
	"errors"
	"reflect"
	"sort"
)

type Iterator struct {
	HasValue     func() bool
	CurrentValue func() Any
	Next         func() Iterator
	//Traits
	Infinite     func() bool
}

type Iterable interface {
	Iterate() Iterator
}

type MappingCallback func(Any) Any

func (it Iterator) Map(mapper MappingCallback) Iterator {
	currentValue := func () Any {
		return mapper(it.CurrentValue())
	}
	next := func () Iterator {
		return it.Next().Map(mapper)
	}
	return Iterator {HasValue: it.HasValue, Next: next, CurrentValue: currentValue}
}

type Predicate func(Any) bool

func (it Iterator) Filter(predicate Predicate) Iterator {
	for it.HasValue() && !predicate(it.CurrentValue()) {
		it = it.Next()
	}

	currentValue := func () Any {
		return it.CurrentValue()
	}
	next := func () Iterator {
		return it.Next().Filter(predicate)
	}

	return Iterator {HasValue: it.HasValue, Next: next, CurrentValue: currentValue}
}

func (it Iterator) Count() int {
	count := 0
	for current := it; current.HasValue(); current = current.Next() {
		count = count + 1
	}
	return count
}

func (it Iterator) Take(count uint64) Iterator {
	currentValue := func () Any {
		return it.CurrentValue()
	}
	next := func () Iterator {
		return it.Next().Take(count - 1)
	}
	hasValue := func () bool {
		return count > 0
	}
	return Iterator {HasValue: hasValue, Next: next, CurrentValue: currentValue}
}

func (it Iterator) Skip(number uint64) Iterator {
	var count uint64 = 0
	current := it
	for ; current.HasValue() && count < number; current = current.Next() {
		count = count + 1
	}
	return current
}

type Reducer func(accum Any, element Any) (Any)

func (it Iterator) Reduce(initialValue Any, reducer Reducer) Any {
	accum := initialValue
	for cursor:=it; cursor.HasValue(); cursor = cursor.Next() {
		accum = reducer(accum, cursor.CurrentValue())
	}
	return accum;
}

func (it Iterator) Reverse() Iterator {
	if it.Infinite != nil && it.Infinite() {
		panic(errors.New("Cannot reverse infinite sequence"))
	}

	list := list.New()
	for current := it; current.HasValue(); current = current.Next() {
		list.PushBack(current.CurrentValue().RawValue())
	}
	return CreateListIterator(list, false)
}

func (it Iterator) ToSlice() []Any {
	accum := make([]Any, 0, 16)
	for current := it; current.HasValue(); current = current.Next() {
		accum = append(accum, current.CurrentValue())
	}
	return accum
}

func IdentityMapping(v Any) Any {
	return v
}

func (it Iterator) ToTypedSliceM(t reflect.Type, converter MappingCallback) interface{} {
	sliceType := reflect.SliceOf(t)
	sliceValue := reflect.MakeSlice(sliceType, 0, 16)
	for current := it; current.HasValue(); current = current.Next() {
		finalElement := reflect.ValueOf(converter(current.CurrentValue()).RawValue())

		sliceValue = reflect.Append(sliceValue, finalElement)
	}
	return sliceValue.Interface()
}

func (it Iterator) ToTypedSlice(t reflect.Type) interface{} {
	return it.ToTypedSliceM(t, IdentityMapping)
}


type ByAny struct {
	slice []Any
	less  func(a, b Any) bool
}

func (ar ByAny) Len() int           { return len(ar.slice) }
func (ar ByAny) Swap(i, j int)      { ar.slice[i], ar.slice[j] = ar.slice[j], ar.slice[i] }
func (ar ByAny) Less(i, j int) bool { return ar.less(ar.slice[i], ar.slice[j]) }

type LessCallback func(a, b Any) bool

func (it Iterator) Sort(less LessCallback) Iterator {
	slice := it.ToSlice()
	controller := ByAny{slice: slice, less: less}
	sort.Sort(controller)
	return CreateArrayIterator(slice)
}

package iterate

import "reflect"

type Iterator struct {
	HasValue     func() bool
	CurrentValue func() Any
	Next         func() Iterator
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

func (it Iterator) ToSlice() []Any {
	accum := make([]Any, 0, 16)
	for current := it; current.HasValue(); current = current.Next() {
		accum = append(accum, current.CurrentValue())
	}
	return accum
}

func (it Iterator) ToTypedSlice(t reflect.Type, converter MappingCallback) interface{} {
	sliceType := reflect.SliceOf(t)
	sliceValue := reflect.MakeSlice(sliceType, 0, 16)
	for current := it; current.HasValue(); current = current.Next() {
		finalElement := reflect.ValueOf(converter(current.CurrentValue()).RawValue())

		sliceValue = reflect.Append(sliceValue, finalElement)
	}
	return sliceValue.Interface()
}


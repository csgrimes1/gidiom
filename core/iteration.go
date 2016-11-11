package core

type Iterator struct {
	HasValue     func() bool
	CurrentValue func() Any
	Next         func() Iterator
}

type Iterable interface {
	Iterate() Iterator
}

func (it Iterator) Map(mapper func(Any) Any) Iterator {
	currentValue := func () Any {
		return mapper(it.CurrentValue())
	}
	next := func () Iterator {
		return it.Next().Map(mapper)
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

func (it Iterator) Skip(number uint64) Iterator {
	var count uint64 = 0
	current := it
	for ; current.HasValue() && count < number; current = current.Next() {
		count = count + 1
	}
	return current
}

func (it Iterator) ToSlice() []Any {
	accum := make([]Any, 0, 16)
	for current := it; current.HasValue(); current = current.Next() {
		accum = append(accum, current.CurrentValue())
	}
	return accum
}

package iterate

func CreateArrayIterator(ar []Any) Iterator {
	hasValue := func () bool {
		return len(ar) >= 1;
	}
	currentValue := func () Any {
		return ar[0]
	}
	next := func () Iterator {
		return CreateArrayIterator(ar[1:])
	}

	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

package iterate

import (
	"container/list"
)

func createForwardIterator(el *list.Element) Iterator {
	hasValue := func () bool {
		return el != nil
	}
	currentValue := func () Any {
		return CreateAny(el.Value)
	}
	next := func () Iterator {
		return createForwardIterator(el.Next())
	}

	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

func createReverseIterator(el *list.Element) Iterator {
	hasValue := func () bool {
		return el != nil
	}
	currentValue := func () Any {
		return CreateAny(el.Value)
	}
	next := func () Iterator {
		return createReverseIterator(el.Prev())
	}

	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

func CreateListIterator(list *list.List, forward bool) Iterator {
	if forward {
		return createForwardIterator(list.Front())
	} else {
		return createReverseIterator(list.Back())
	}
}

package core

type Sequencer func(context Any) (Any, Any)

func IterateSequence(generator Sequencer, startingContext Any) Iterator {
	context, val := generator(startingContext)

	hasValue := func () bool {
		return true
	}
	currentValue := func () Any {
		return val
	}
	next := func () Iterator {
		return IterateSequence(generator, context)
	}

	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

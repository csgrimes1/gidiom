package iterate

type Sequencer func(context Any) (Any, Any)

func CreateSequence(startingContext Any, generator Sequencer) Iterator {
	context, val := generator(startingContext)

	hasValue := func () bool {
		return true
	}
	currentValue := func () Any {
		return val
	}
	next := func () Iterator {
		return CreateSequence(context, generator)
	}

	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

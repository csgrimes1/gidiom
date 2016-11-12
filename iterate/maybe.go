package iterate

type Maybe struct {
	value []Any
}

func CreateMaybe(value *Any) Maybe {
	if value == nil {
		ar := []Any{}
		return Maybe{value: ar}
	} else {
		ar := []Any{*value}
		return Maybe{value: ar}
	}
	return Maybe{}
}

func (self Maybe) HasValue() bool {
	return len(self.value) > 0
}

func (self Maybe) Value() Any {
	return self.value[0]
}

func (self Maybe) Iterate() Iterator {
	return CreateArrayIterator(self.value)
}


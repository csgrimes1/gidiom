package iterate

type Any struct {
	value interface{}
}

func CreateAny(value interface{}) Any {
	return Any{value: value}
}

func (any Any) RawValue() interface{} {
	return any.value
}

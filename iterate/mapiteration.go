package iterate

import "reflect"

type MapPair struct {
	key reflect.Value
	value reflect.Value
}
var MapPairType reflect.Type = reflect.TypeOf(MapPair{})

func (mp MapPair) Key() Any {
	return CreateAny(mp.key.Interface())
}

func (mp MapPair) Value() Any {
	return CreateAny(mp.value.Interface())
}

func iterateOverKeys(keys []reflect.Value, getter func(key reflect.Value) MapPair) Iterator {
	hasValue := func () bool {
		return len(keys) >= 1;
	}
	currentValue := func () Any {
		key := keys[0]
		return CreateAny(getter(key))
	}
	next := func () Iterator {
		return iterateOverKeys(keys[1:], getter)
	}
	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

func CreateMapIterator(theMap interface{}) Iterator {
	m := reflect.ValueOf(theMap)
	getter := func(key reflect.Value) MapPair {
		return MapPair{key: key,
			value: m.MapIndex(key)}
	}

	return iterateOverKeys(m.MapKeys(), getter)
}

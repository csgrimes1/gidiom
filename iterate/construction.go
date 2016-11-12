package iterate

import (
	"fmt"
	"reflect"
)

func forArray(ar reflect.Value) Iterator {
	hasValue := func () bool {
		return ar.Len() >= 1;
	}
	currentValue := func () Any {
		return CreateAny(ar.Index(0).Interface())
	}
	next := func () Iterator {
		return forArray(ar.Slice(1, ar.Len()))
	}

	return Iterator {HasValue: hasValue, CurrentValue: currentValue, Next: next}
}

func Start(data interface{}) Iterator {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Slice:
		return forArray(v)
	case reflect.Map:
		fmt.Printf("map: %v\n", v.Interface())
	default:
		any := CreateAny(data)
		maybe := CreateMaybe(&any)
		return maybe.Iterate()
	}

	return Iterator{}
}

func MAP(mappingFunc interface{}) MappingCallback {
	fv := reflect.ValueOf(mappingFunc)

	return func(input Any) Any {
		inputValue := reflect.ValueOf(input.RawValue())
		inputs := []reflect.Value{inputValue}
		results := fv.Call(inputs)
		return CreateAny(results[0].Interface())
	}
}

func PRED(predicate interface{}) Predicate {
	fv := reflect.ValueOf(predicate)

	return func(input Any) bool {
		inputValue := reflect.ValueOf(input.RawValue())
		inputs := []reflect.Value{inputValue}
		results := fv.Call(inputs)
		return results[0].Interface().(bool)
	}
}

func GENERATOR(generator interface{}) Sequencer {
	fv := reflect.ValueOf(generator)

	return func(context Any) (Any, Any) {
		inputValue := reflect.ValueOf(context.RawValue())
		inputs := []reflect.Value{inputValue}
		results := fv.Call(inputs)
		return CreateAny(results[0].Interface()), CreateAny(results[1].Interface())
	}
}
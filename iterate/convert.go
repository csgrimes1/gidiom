package iterate

import (
	"reflect"
	"strconv"
)

var int64Type = reflect.TypeOf(int64(1))
var uint64Type = reflect.TypeOf(uint64(1))
var float64Type = reflect.TypeOf(float64(1.0))
var boolType = reflect.TypeOf(false)

type Converter struct {
	value interface{}
}

func createConverter(value interface{}) Converter {
	return Converter{value: value}
}

func (c Converter) Int64() int64 {
	var err interface{}
	var i int64
	switch x := c.value.(type) {
	case string:
		i, err = strconv.ParseInt(x, 10, 64)
	default:
		i = reflect.ValueOf(c.value).Convert(int64Type).Interface().(int64)
	}
	if err != nil {
		panic(err)
	}
	return i
}

func (c Converter) Int32() int32 {
	return int32(c.Int64())
}

func (c Converter) Int16() int16 {
	return int16(c.Int64())
}

func (c Converter) Int8() int8 {
	return int8(c.Int64())
}

func (c Converter) Int() int {
	return int(c.Int64())
}


func (c Converter) Uint64() uint64 {
	var err interface{}
	var i uint64
	switch x := c.value.(type) {
	case string:
		i, err = strconv.ParseUint(x, 10, 64)
	default:
		i = reflect.ValueOf(c.value).Convert(uint64Type).Interface().(uint64)
	}
	if err != nil {
		panic(err)
	}
	return i
}

func (c Converter) Uint32() uint32 {
	return uint32(c.Uint64())
}

func (c Converter) Uint16() uint16 {
	return uint16(c.Uint64())
}

func (c Converter) Uint8() uint8 {
	return uint8(c.Uint64())
}

func (c Converter) Uint() uint {
	return uint(c.Uint64())
}

func (c Converter) Float64() float64 {
	var err interface{}
	var f float64
	switch x := c.value.(type) {
	case string:
		f, err = strconv.ParseFloat(x, 64)
	default:
		f = reflect.ValueOf(c.value).Convert(float64Type).Interface().(float64)
	}
	if err != nil {
		panic(err)
	}
	return f
}

func (c Converter) Float32() float32 {
	return float32(c.Float64())
}

func (c Converter) Bool() bool {
	var err interface{}
	var b bool
	switch x := c.value.(type) {
	case string:
		b, err = strconv.ParseBool(x)
	case int64:
		b = (x != 0)
	case int32:
		b = (x != 0)
	case int16:
		b = (x != 0)
	case int8:
		b = (x != 0)
	case int:
		b = (x != 0)
	case uint64:
		b = (x != 0)
	case uint32:
		b = (x != 0)
	case uint16:
		b = (x != 0)
	case uint8:
		b = (x != 0)
	case uint:
		b = (x != 0)
	case float64:
		b = (x != 0.0)
	case float32:
		b = (x != 0.0)
	default:
		b = reflect.ValueOf(c.value).Convert(boolType).Interface().(bool)
	}
	if err != nil {
		panic(err)
	}
	return b
}

func (c Converter) String() string {
	switch x := c.value.(type) {
	case string:
		return x
	case int64:
		return strconv.FormatInt(x, 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case int:
		return strconv.FormatInt(int64(x), 10)
	case float64:
		return strconv.FormatFloat(x, 'g', 5, 64)
	case float32:
		return strconv.FormatFloat(float64(x), 'g', 5, 32)
	case uint64:
		return strconv.FormatUint(x, 10)
	case uint32:
		return strconv.FormatUint(uint64(x), 10)
	case uint16:
		return strconv.FormatUint(uint64(x), 10)
	case uint8:
		return strconv.FormatUint(uint64(x), 10)
	case uint:
		return strconv.FormatUint(uint64(x), 10)
	case bool:
		return strconv.FormatBool(x)
	}

	return ""
}

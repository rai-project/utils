package utils

import (
	"reflect"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

func FlattenBoolSlice(data interface{}) []bool {
	// when data is a scalar
	rval := reflect.ValueOf(data)
	typ := rval.Type()
	if typ.Kind() != reflect.Array &&
		typ.Kind() != reflect.Slice &&
		typ.Kind() != reflect.Interface {
		ddata := indirect(data)
		if e, ok := ddata.(bool); ok {
			return []bool{e}
		}
		switch s := ddata.(type) {
		case bool:
			return []bool{bool(s)}
		case int:
			return []bool{bool(s == 1)}
		case uint:
			return []bool{bool(s == 1)}
		case int8:
			return []bool{bool(s == 1)}
		case uint8:
			return []bool{bool(s == 1)}
		case int16:
			return []bool{bool(s == 1)}
		case uint16:
			return []bool{bool(s == 1)}
		case int32:
			return []bool{bool(s == 1)}
		case uint32:
			return []bool{bool(s == 1)}
		case int64:
			return []bool{bool(s == 1)}
		case uint64:
			return []bool{bool(s == 1)}
		case float32:
			return []bool{bool(s == 1)}
		case float64:
			return []bool{bool(s == 1)}
		case uintptr:
			return []bool{bool(s == 1)}
		}
		panic(errors.Errorf("unable to convert %v of kind %v", pp.Sprint(data), typ.Kind().String()))
	}

	// no we know data is a slice
	res := []bool{}
	for ii := 0; ii < rval.Len(); ii++ {
		val := rval.Index(ii)
		fval := FlattenBoolSlice(val.Interface())
		res = append(res, fval...)
	}
	return res
}

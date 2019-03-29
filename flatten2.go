//go:generate go get github.com/cheekybits/genny
//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ElementType=int,uint,uintptr,uint8,uint16,uint32,uint64,int8,int16,int32,int64,float32,float64"

package utils

import (
	"reflect"

	"github.com/k0kubun/pp"

	"github.com/cheekybits/genny/generic"
	"github.com/pkg/errors"
)

type ElementType generic.Type

func FlattenElementTypeSlice(data interface{}) []ElementType {
	// when data is a scalar
	rval := reflect.ValueOf(data)
	typ := rval.Type()
	if typ.Kind() != reflect.Array &&
		typ.Kind() != reflect.Slice &&
		typ.Kind() != reflect.Interface {
		ddata := indirect(data)
		if e, ok := ddata.(ElementType); ok {
			return []ElementType{e}
		}
		switch s := ddata.(type) {
		case bool:
			if s {
				return []ElementType{ElementType(1)}
			}
			return []ElementType{ElementType(0)}
		case int:
			return []ElementType{ElementType(s)}
		case uint:
			return []ElementType{ElementType(s)}
		case int8:
			return []ElementType{ElementType(s)}
		case uint8:
			return []ElementType{ElementType(s)}
		case int16:
			return []ElementType{ElementType(s)}
		case uint16:
			return []ElementType{ElementType(s)}
		case int32:
			return []ElementType{ElementType(s)}
		case uint32:
			return []ElementType{ElementType(s)}
		case int64:
			return []ElementType{ElementType(s)}
		case uint64:
			return []ElementType{ElementType(s)}
		case float32:
			return []ElementType{ElementType(s)}
		case float64:
			return []ElementType{ElementType(s)}
		case uintptr:
			return []ElementType{ElementType(s)}
		}
		panic(errors.Errorf("unable to convert %v of kind %v", pp.Sprint(data), typ.Kind().String()))
	}

	// no we know data is a slice
	res := []ElementType{}
	for ii := 0; ii < rval.Len(); ii++ {
		val := rval.Index(ii)
		fval := FlattenElementTypeSlice(val.Interface())
		res = append(res, fval...)
	}
	return res
}

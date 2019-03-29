package utils

import "github.com/tinygo-org/tinygo/src/reflect"

func numElements(shape []int64) int64 {
	n := int64(1)
	for _, d := range shape {
		n *= d
	}
	return n
}

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func shapeAndTypeOf(val reflect.Value) (shape []int64, dt reflect.Type, err error) {
	typ := val.Type()
	for typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
		shape = append(shape, int64(val.Len()))
		if val.Len() > 0 {
			// In order to check tensor structure properly in general case we need to iterate over all slices of the tensor to check sizes match
			// Since we already going to iterate over all elements in encodeTensor() let's
			// 1) do the actual check in encodeTensor() to save some cpu cycles here
			// 2) assume the shape is represented by lengths of elements with zero index in each dimension
			val = val.Index(0)
		}
		typ = typ.Elem()
	}
	return shape, typ, nil
}

func ShapeAndTypeOf(val interface{}) (shape []int64, dt reflect.Type, err error) {
	return shapeAndTypeOf(reflect.ValueOf(val))
}

func FlattenedLength(val interface{}) (int64, error) {
	shape, _, err := ShapeAndTypeOf(val)
	if err != nil {
		return int64(0), err
	}
	return numElements(shape), nil
}

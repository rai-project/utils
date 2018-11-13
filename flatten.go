package utils

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

// casts an interface to a []float32 type.
func Tofloat32SliceE(i interface{}) ([]float32, error) {
	if i == nil {
		return []float32{}, fmt.Errorf("unable to cast %#v of type %T to []float32", i, i)
	}

	switch v := i.(type) {
	case []float32:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]float32, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := cast.Tofloat32E(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []float32", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []float32", i, i)
	}
}

// func FlattenFloat32(f interface{}) []float32 {

// 	output := []float32{}

// 	switch e := f.(type) {
// 	case [][]float32:
// 		for _, v := range e {
// 			output = append(output, v...)
// 		}
// 	case []interface{}:
// 		for _, v := range e {
// 			output = append(output, FlattenFloat32(v)...)
// 		}
// 	case float32:
// 		output = append(output, e)
// 	default:
// 		output = append(output, cast.ToFloat32(e))
// 	}
// 	return output
// }

func Flatten(f interface{}) []interface{} {

	output := []interface{}{}

	switch e := f.(type) {
	case []interface{}:
		for _, v := range e {
			output = append(output, Flatten(v)...)
		}
	case interface{}:
		output = append(output, e)
	}

	return output
}

func Flatten2DFloat32(f interface{}) [][]float32 {

	switch e := f.(type) {
	case [][][][]float32:
		output := [][]float32{}
		for _, v := range e {
			output = append(output, Flatten2DFloat32(v)...)
		}
		return output
	case [][][]float32:
		output := [][]float32{}
		for _, v := range e {
			output = append(output, Flatten2DFloat32(v)...)
		}
		return output
	case [][]float32:
		return e
	case []float32:
		return [][]float32{e}
	case []interface{}:
		output := [][]float32{}
		for _, v := range e {
			output = append(output, Flatten2DFloat32(v)...)
		}
		return output
	default:
		panic("expecting a float value while flattening float32...")
	}

	return nil
}

package utils

import (
	"github.com/spf13/cast"
)

func FlattenFloat32(f interface{}) []float32 {

	output := []float32{}

	switch e := f.(type) {
	case [][]float32:
		for _, v := range e {
			output = append(output, v...)
		}
	case []interface{}:
		for _, v := range e {
			output = append(output, FlattenFloat32(v)...)
		}
	case float32:
		output = append(output, e)
	default:
		output = append(output, cast.ToFloat32(e))
	}
	return output
}

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

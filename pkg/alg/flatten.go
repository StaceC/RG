package alg

import (
	"reflect"
)

// https://www.golangprograms.com/go-language/arrays.html
// Flatten - an arbitrarily nested array of values
func Flatten(input []interface{}) []interface{} {

	out := make([]interface{}, 0)

	for _, v := range input {
		if reflect.ValueOf(v).Kind() == reflect.Array ||
			reflect.ValueOf(v).Kind() == reflect.Slice {
			out = append(out, Flatten(v.([]interface{}))...)
		} else {
			out = append(out, v)
		}
	}

	return out
}

// Package structs provides a set of functions to manipulate structs.
//
// Example
//
//	import "github.com/neurocode-io/go-pkgs/structs"
//
//	input := struct {
//		Name string
//		Age  int
//	}{
//		Name: "John",
//		Age:  30,
//	}
//
//	result := structs.ToMap(input)
//	// result == map[string]any{"Name": "John", "Age": 30}
package structs

import (
	"fmt"
	"reflect"
)

/*
ToMap converts a struct to a map[string]any.

Example

	input := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}
	result := structs.ToMap(input)
	// result == map[string]any{"Name": "John", "Age": 30}
*/
//nolint:gocognit,cyclop
func ToMap(input any) any {
	val := reflect.ValueOf(input)

	switch val.Kind() { //nolint:exhaustive
	case reflect.Pointer:
		// Dereference pointers
		if val.IsNil() {
			return nil
		}

		return ToMap(val.Elem().Interface())

	case reflect.Struct:
		// Convert struct to map
		out := make(map[string]any)
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			// Skip unexported fields
			if field.PkgPath != "" {
				continue
			}
			out[field.Name] = ToMap(val.Field(i).Interface())
		}

		return out

	case reflect.Slice, reflect.Array:
		// Convert slices/arrays to a slice of interfaces
		length := val.Len()
		out := make([]any, 0, length)
		for i := 0; i < length; i++ {
			out = append(out, ToMap(val.Index(i).Interface()))
		}

		return out

	case reflect.Map:
		// Convert maps to map[string]any
		out := make(map[string]any)
		for _, key := range val.MapKeys() {
			// We assume string keys; otherwise convert them as needed
			strKey := fmt.Sprintf("%v", key.Interface())
			out[strKey] = ToMap(val.MapIndex(key).Interface())
		}

		return out

	default:
		// basic types, return as-is
		return input
	}
}

// Package slice provides a set of functions to manipulate slices.
//
// Example:
//
//	import "github.com/neurocode-io/go-pkgs/pkg/slice"
//
//	input := []int{1, 2, 3}
//	slice.Reverse(input)
//	// input == []int{3, 2, 1}
package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type (
	mapFunc[T any]  func(T) T
	keepFunc[T any] func(T) bool
)

// Contains returns true if the slice contains the value.
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}

	return false
}

// Reverse returns a reversed copy of the slice.
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return result
}

// Sort returns a sorted copy of the slice.
func Sort[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

// Map returns a new slice with the result of applying the function to each element.
// Example:
//
//		input := []int{1, 2, 3}
//	 fn := func(i int) int { return i * 2 }
//		result := slice.Map(input, fn)
//		// result == []int{2, 4, 6}
func Map[T any](slice []T, fn mapFunc[T]) []T {
	result := make([]T, len(slice))
	for i := range slice {
		result[i] = fn(slice[i])
	}

	return result
}

// Reduce reduces the slice to a single value.
// The function is called with an initializer sometimes called an accumulator and each element in the slice.
// The result of the function is passed as the accumulator to the next call.
// Example:
//
//	input := []int{1, 2, 3}
//	fn := func(acc int, i int) int { return acc + i }
//	result := slice.Reduce(input, 0, fn)
//	// result == 6
func Reduce[T, V any](slice []T, accumulator V, fn func(V, T) V) V {
	result := accumulator
	for _, value := range slice {
		result = fn(result, value)
	}

	return result
}

// Merge merges multiple slices into a single slice.
// Example:
//
//	input1 := []int{1, 2, 3}
//	input2 := []int{4, 5, 6}
//	result := slice.Merge(input1, input2)
//	// result == []int{1, 2, 3, 4, 5, 6}
func Merge[T any](slices ...[]T) []T {
	result := []T{}
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return result
}

// Filter returns a new slice with the elements that match the function.
// Example:
//
//	input := []int{1, 2, 3, 4, 5}
//	fn := func(i int) bool { return i%2 == 0 }
//	result := slice.Filter(input, fn)
//	// result == []int{2, 4}
func Filter[T any](slice []T, fn keepFunc[T]) []T {
	result := make([]T, len(slice))
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Any returns true if any element matches the function.
// Example:
//
//	input := []string{"a", "b", "c"}
//	fn := func(s string) bool { return s == "b" }
//	slice.Any(input, fn)
//	// true
func Any[T any](slice []T, fn keepFunc[T]) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}

	return false
}

// All returns true if all elements match the function.
// Example:
//
//	input := []string{"a", "b", "c"}
//	fn := func(s string) bool { return len(s) == 1 }
//	slice.All(input, fn)
//	// true
func All[T any](slice []T, fn keepFunc[T]) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}

	return true
}

// None returns true if no elements match the function.
// Example:
//
//	input := []string{"a", "b", "c"}
//	fn := func(s string) bool { return len(s) == 2 }
//	slice.None(input, fn)
//	// true
func None[T any](slice []T, fn keepFunc[T]) bool {
	return !Any(slice, fn)
}

// Unique returns a new slice with unique elements.
// Example:
//
//	input := []string{"a", "b", "c", "c", "b", "a"}
//	slice.Unique(input)
//	// input == []string{"a", "b", "c"}
func Unique[T comparable](slice []T) []T {
	result := []T{}
	for _, v := range slice {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}

	return result
}

// UniqueBy returns a new slice with unique elements based on the function.
// Example:
//
//	input := []string{"a", "b", "c", "aa", "bb", "cc"}
//	fn := func(s string) string {
//		return s[:1]
//	}
//	slice.UniqueBy(input, fn)
//	// input == []string{"a", "b", "c"}
func UniqueBy[T comparable](slice []T, fn mapFunc[T]) []T {
	result := []T{}
	for _, v := range slice {
		if !Contains(Map(result, fn), fn(v)) {
			result = append(result, v)
		}
	}

	return result
}

// Find returns the first element that matches the function.
// If no element matches the function, the zero value is returned.
// Example:
//
//	input := []string{"a", "b", "c"}
//	fn := func(s string) bool {
//		return s == "b"
//	}
//	slice.Find(input, fn)
//	// input == "b"
func Find[T comparable](slice []T, fn keepFunc[T]) (T, bool) {
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	var result T

	return result, false
}

// FindIndex returns the index of the first element that matches the function.
// If no element matches the function, -1 is returned.
// Example:
//
//	input := []string{"a", "b", "c"}
//
//	fn := func(s string) bool {
//			return s == "b"
//	}
//
//	slice.FindIndex(input, fn)
//	// input == 1
func FindIndex[T comparable](slice []T, fn keepFunc[T]) (int, bool) {
	for i, v := range slice {
		if fn(v) {
			return i, true
		}
	}

	return -1, false
}

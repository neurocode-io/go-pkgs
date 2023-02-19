// Package m provides functions for working with maps.
//
// Example:
//
//	import m "github.com/neurocode-io/go-pkgs/pkg/map"
//
//	input := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := m.Keys(input)

package m

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Keys returns a slice of the keys in the map.
// Example:
//
//	input := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := m.Keys(input)
//	// keys == []string{"a", "b", "c"}
func Keys[T comparable, V any](input map[T]V) []T {
	result := make([]T, 0, len(input))
	for i := range input {
		result = append(result, i)
	}

	return result
}

// Sum returns the sum of the values in the map.
// Example:
//
//	input := map[string]int{"a": 1, "b": 2, "c": 3}
//	sum := m.Sum(input)
//	// sum == 6
func Sum[T comparable, V constraints.Float | constraints.Integer](input map[T]V) V {
	var result V
	for _, v := range input {
		result += v
	}

	return result
}

// Values returns a slice of the values in the map.
// Example:
//
//	input := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := m.Values(input)
//	// values == []int{1, 2, 3}
func Values[T comparable, V any](input map[T]V) []V {
	result := make([]V, 0, len(input))
	for _, v := range input {
		result = append(result, v)
	}

	return result
}

// Entries returns a slice of the entries in the map.
// Example:
//
//	input := map[string]int{"a": 1, "b": 2, "c": 3}
//	entries := m.Entries(input)
//	// entries == [][2]string{{"a", 1}, {"b", 2}, {"c", 3}}
func Entries[T comparable](input map[T]T) [][2]T {
	result := make([][2]T, 0, len(input))
	for i, j := range input {
		result = append(result, [2]T{i, j})
	}

	return result
}

// FromEntries returns a map from the entries in the slice.
// Example:
//
//	input := [][2]string{{"a", 1}, {"b", 2}, {"c", 3}}
//	m := m.FromEntries(input)
//	// m == map[string]int{"a": 1, "b": 2, "c": 3}
func FromEntries[T comparable](entries [][2]T) map[T]T {
	result := make(map[T]T)
	for _, entry := range entries {
		result[entry[0]] = entry[1]
	}

	return result
}

// SortByValue returns a map sorted by the values.
// Example:
//
//	input := map[string]int{"a": 9, "b": 1, "c": 3}
//	m := m.SortByValue(input)
//	// m == map[string]int{"b": 1, "c": 3, "a": 9}
func SortByValue[T constraints.Ordered](input map[T]T) map[T]T {
	entries := Entries(input)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i][1] < entries[j][1]
	})

	return FromEntries(entries)
}

// Sort returns a map sorted by the keys.
// Example:
//
//	input := map[string]string{"b": "world", "a": "hello", "c": "!"}
//	m := m.Sort(input)
//	// m == map[string]string{"a": "hello", "b": "world", "c": "!"}
func Sort[T constraints.Ordered](input map[T]T) map[T]T {
	entries := Entries(input)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i][0] < entries[j][0]
	})

	return FromEntries(entries)
}

package m

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Keys[T comparable, V any](input map[T]V) []T {
	result := make([]T, 0, len(input))
	for i := range input {
		result = append(result, i)
	}

	return result
}

func Sum[T comparable, V constraints.Float | constraints.Integer](input map[T]V) V {
	var result V
	for _, v := range input {
		result += v
	}

	return result
}

func Values[T comparable, V any](input map[T]V) []V {
	result := make([]V, 0, len(input))
	for _, v := range input {
		result = append(result, v)
	}

	return result
}

func Entries[T comparable](input map[T]T) [][2]T {
	result := make([][2]T, 0, len(input))
	for i, j := range input {
		result = append(result, [2]T{i, j})
	}

	return result
}

func FromEntries[T comparable](entries [][2]T) map[T]T {
	result := make(map[T]T)
	for _, entry := range entries {
		result[entry[0]] = entry[1]
	}

	return result
}

func SortByValue[T constraints.Ordered](input map[T]T) map[T]T {
	entries := Entries(input)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i][1] < entries[j][1]
	})

	return FromEntries(entries)
}

func Sort[T constraints.Ordered](input map[T]T) map[T]T {
	entries := Entries(input)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i][0] < entries[j][0]
	})

	return FromEntries(entries)
}

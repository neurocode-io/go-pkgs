package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type (
	mapFunc[T any]  func(T) T
	keepFunc[T any] func(T) bool
)

func Contains[T comparable](s []T, v T) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}

	return false
}

func Reverse[T any](s []T) []T {
	result := make([]T, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return result
}

func Sort[T constraints.Ordered](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func Map[T any](s []T, f mapFunc[T]) []T {
	result := make([]T, len(s))
	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

func Reduce[T, V any](s []T, initializer V, f func(V, T) V) V {
	result := initializer
	for _, value := range s {
		result = f(result, value)
	}

	return result
}

func Merge[T any](slices ...[]T) []T {
	result := []T{}
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return result
}

func Filter[T any](s []T, f keepFunc[T]) []T {
	result := []T{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Any[T any](s []T, f keepFunc[T]) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}

	return false
}

func All[T any](s []T, f keepFunc[T]) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

func None[T any](s []T, f keepFunc[T]) bool {
	return !Any(s, f)
}

func Unique[T comparable](s []T) []T {
	result := []T{}
	for _, v := range s {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}

	return result
}

func UniqueBy[T comparable](s []T, f mapFunc[T]) []T {
	result := []T{}
	for _, v := range s {
		if !Contains(Map(result, f), f(v)) {
			result = append(result, v)
		}
	}

	return result
}

func Find[T comparable](s []T, f keepFunc[T]) (T, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	var result T

	return result, false
}

func FindIndex[T comparable](s []T, f keepFunc[T]) (int, bool) {
	for i, v := range s {
		if f(v) {
			return i, true
		}
	}

	return -1, false
}

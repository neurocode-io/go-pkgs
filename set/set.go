// Package set provides a generic set implementation.
package set

// Set is a generic set implementation.
type Set[T comparable] map[T]struct{}

// New returns a new set.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

/*
FromSlice returns a new set from a slice.

Example

	s := FromSlice([]int{1, 2, 3, 3, 3, 4})
	s.ToSlice() // []int{1, 2, 3, 4}
*/
func FromSlice[T comparable](s []T) Set[T] {
	result := New[T]()
	for _, v := range s {
		result.Add(v)
	}

	return result
}

/*
Add adds a value to the set.

Example

	s := New[int]()
	s.Add(1)
	s.Add(2)
*/
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

/*
Remove removes a value from the set.

Example

	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Remove(1)
*/
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

/*
Exists returns true if the value exists in the set.

Example

	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Exists(1) // true
	s.Exists(3) // false
*/
func (s Set[T]) Exists(v T) bool {
	_, ok := s[v]

	return ok
}

/*
Size returns the number of elements in the set.

Example

	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Size() // 2
*/
func (s Set[T]) Size() int {
	return len(s)
}

/*
ToSlice returns a slice of the values in the set.

Example

	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.ToSlice() // []int{1, 2}
*/
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}

	return result
}

/*
Union returns a new set that is the union of the two sets.

Example

	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s1.Union(s2)
	s3.ToSlice() // []int{1, 2, 3}
*/
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s {
		result.Add(v)
	}
	for v := range other {
		result.Add(v)
	}

	return result
}

/*
Intersection returns a new set that is the intersection of the two sets.

Example

	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s1.Intersection(s2)
	s3.ToSlice() // []int{2}
*/
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s {
		if other.Exists(v) {
			result.Add(v)
		}
	}

	return result
}

/*
Difference returns a new set that is the difference of the two sets.

Example

	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s1.Difference(s2)
	s3.ToSlice() // []int{1}
*/
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s {
		if !other.Exists(v) {
			result.Add(v)
		}
	}

	return result
}

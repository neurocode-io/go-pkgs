package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	t.Run("FromSlice", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3, 3, 3, 4})

		assert.True(t, s.Exists(1))
		assert.True(t, s.Exists(2))
		assert.True(t, s.Exists(3))
		assert.True(t, s.Exists(4))
		assert.Equal(t, 4, s.Size())
		assert.ElementsMatch(t, []int{1, 2, 3, 4}, s.ToSlice())
	})

	t.Run("Add", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		s.Add(2)

		assert.True(t, s.Exists(1))
		assert.True(t, s.Exists(2))
		assert.Equal(t, 2, s.Size())
		assert.ElementsMatch(t, []int{1, 2}, s.ToSlice())
	})

	t.Run("Remove", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		s.Add(2)
		s.Remove(1)

		assert.False(t, s.Exists(1))
		assert.True(t, s.Exists(2))
		assert.Equal(t, 1, s.Size())
		assert.Equal(t, []int{2}, s.ToSlice())
	})

	t.Run("Union", func(t *testing.T) {
		s := New[int]()
		s.Add(2)

		s2 := New[int]()
		s2.Add(3)
		s2.Add(4)

		s3 := s.Union(s2)

		assert.True(t, s3.Exists(2))
		assert.True(t, s3.Exists(3))
		assert.True(t, s3.Exists(4))
		assert.Equal(t, 3, s3.Size())
		assert.ElementsMatch(t, []int{2, 3, 4}, s3.ToSlice())
	})

	t.Run("Intersection", func(t *testing.T) {
		s := New[int]()
		s.Add(2)
		s.Add(3)

		s2 := New[int]()
		s2.Add(3)
		s2.Add(4)

		s3 := s.Intersection(s2)

		assert.False(t, s3.Exists(2))
		assert.True(t, s3.Exists(3))
		assert.False(t, s3.Exists(4))
		assert.Equal(t, 1, s3.Size())
		assert.ElementsMatch(t, []int{3}, s3.ToSlice())
	})

	t.Run("Difference", func(t *testing.T) {
		s := New[int]()
		s.Add(2)
		s.Add(3)

		s2 := New[int]()
		s2.Add(3)
		s2.Add(4)

		s3 := s.Difference(s2)

		assert.True(t, s3.Exists(2))
		assert.False(t, s3.Exists(3))
		assert.False(t, s3.Exists(4))
		assert.Equal(t, 1, s3.Size())
		assert.ElementsMatch(t, []int{2}, s3.ToSlice())
	})
}

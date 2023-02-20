package m

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsyncMap(t *testing.T) {
	t.Run("Store/Load", func(t *testing.T) {
		m := NewAsyncMap[string, int]()
		m.Store("a", 1)
		v, ok := m.Load("a")

		assert.True(t, ok)
		assert.Equal(t, 1, v)
	})

	t.Run("LoadOrStore", func(t *testing.T) {
		m := NewAsyncMap[string, int]()
		v, ok := m.LoadOrStore("a", 1)

		assert.False(t, ok)
		assert.Equal(t, 1, v)

		v, ok = m.LoadOrStore("a", 2)

		assert.True(t, ok)
		assert.Equal(t, 1, v)
	})

	t.Run("LoadAndDelete", func(t *testing.T) {
		m := NewAsyncMap[string, int]()
		m.Store("a", 1)
		v, ok := m.LoadAndDelete("a")

		assert.True(t, ok)
		assert.Equal(t, 1, v)

		v, ok = m.LoadAndDelete("a")

		assert.False(t, ok)
		assert.Equal(t, 0, v)
	})

	t.Run("Range", func(t *testing.T) {
		m := NewAsyncMap[string, int]()
		m.Store("a", 1)
		m.Store("b", 2)
		runs := 0
		fn := func(key string, value int) bool {
			runs++
			if key == "a" {
				assert.Equal(t, 1, value)
			}

			if key == "b" {
				assert.Equal(t, 2, value)
			}

			return true
		}
		m.Range(fn)
		assert.Equal(t, 2, runs)

		runs = 0
		fn2 := func(key string, value int) bool {
			runs++
			return false
		}

		m.Range(fn2)
		assert.Equal(t, 1, runs)
	})

	t.Run("Delete", func(t *testing.T) {
		m := NewAsyncMap[string, int]()
		m.Store("a", 1)
		m.Delete("a")
		v, ok := m.Load("a")

		assert.False(t, ok)
		assert.Equal(t, 0, v)
	})
}

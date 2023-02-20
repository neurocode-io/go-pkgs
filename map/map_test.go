package m

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		assert.ElementsMatch(t, []string{"a", "b", "c"}, Keys(m))

		m2 := map[int]string{1: "a", 2: "b", 3: "c"}
		assert.ElementsMatch(t, []int{1, 2, 3}, Keys(m2))
	})

	t.Run("Sum", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		assert.Equal(t, 6, Sum(m))

		m2 := map[int]int{1: 1, 2: 2, 3: 3}
		assert.Equal(t, 6, Sum(m2))
	})

	t.Run("Values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		assert.ElementsMatch(t, []int{1, 2, 3}, Values(m))

		m2 := map[int]string{1: "a", 2: "b", 3: "c"}
		assert.ElementsMatch(t, []string{"a", "b", "c"}, Values(m2))
	})

	t.Run("Entries", func(t *testing.T) {
		m := map[string]string{"a": "a", "b": "b", "c": "c"}
		assert.ElementsMatch(t, [][2]string{{"a", "a"}, {"b", "b"}, {"c", "c"}}, Entries(m))

		m2 := map[int]int{1: 1, 2: 2, 3: 3}
		assert.ElementsMatch(t, [][2]int{{1, 1}, {2, 2}, {3, 3}}, Entries(m2))
	})

	t.Run("FromEntries", func(t *testing.T) {
		m := [][2]string{{"a", "a"}, {"b", "b"}, {"c", "c"}}
		assert.Equal(t, map[string]string{"a": "a", "b": "b", "c": "c"}, FromEntries(m))

		m2 := [][2]int{{1, 1}, {2, 2}, {3, 3}}
		assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, FromEntries(m2))
	})

	t.Run("SortByValue", func(t *testing.T) {
		m := map[string]string{"a": "c", "b": "a", "c": "b"}
		assert.Equal(t, map[string]string{"b": "a", "c": "b", "a": "c"}, SortByValue(m))

		m2 := map[int]int{1: 3, 2: 1, 3: 2}
		assert.Equal(t, map[int]int{2: 1, 3: 2, 1: 3}, SortByValue(m2))
	})

	t.Run("Sort", func(t *testing.T) {
		m := map[string]string{"c": "a", "b": "b", "a": "c"}
		assert.Equal(t, map[string]string{"a": "c", "b": "b", "c": "a"}, Sort(m))

		m2 := map[int]int{1: 3, 2: 1, 3: 2}
		assert.Equal(t, map[int]int{1: 3, 2: 1, 3: 2}, Sort(m2))
	})
}

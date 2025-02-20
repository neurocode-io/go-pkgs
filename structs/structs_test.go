package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	t.Run("basic types", func(t *testing.T) {
		t.Run("returns integers as-is", func(t *testing.T) {
			result := ToMap(42)
			assert.Equal(t, 42, result)
		})

		t.Run("returns strings as-is", func(t *testing.T) {
			result := ToMap("hello")
			assert.Equal(t, "hello", result)
		})
	})

	t.Run("pointers", func(t *testing.T) {
		t.Run("handles nil pointer", func(t *testing.T) {
			var ptr *string
			result := ToMap(ptr)
			assert.Nil(t, result)
		})

		t.Run("dereferences pointer", func(t *testing.T) {
			val := "test"
			result := ToMap(&val)
			assert.Equal(t, "test", result)
		})
	})

	t.Run("structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		t.Run("converts struct to map", func(t *testing.T) {
			person := Person{Name: "Alice", Age: 30}
			result := ToMap(person)

			expected := map[string]any{
				"Name": "Alice",
				"Age":  30,
			}
			assert.Equal(t, expected, result)
		})

		t.Run("handles nested structs", func(t *testing.T) {
			type Address struct {
				City    string
				Country string
			}
			type Employee struct {
				Person
				Address Address
			}

			emp := Employee{
				Person:  Person{Name: "Bob", Age: 25},
				Address: Address{City: "Reykjavik", Country: "Iceland"},
			}

			expected := map[string]any{
				"Person": map[string]any{
					"Name": "Bob",
					"Age":  25,
				},
				"Address": map[string]any{
					"City":    "Reykjavik",
					"Country": "Iceland",
				},
			}
			assert.Equal(t, expected, ToMap(emp))
		})

		t.Run("handles slices", func(t *testing.T) {
			type Person struct {
				Name string
				Age  int
			}

			type Company struct {
				Employees []Person
			}

			company := Company{
				Employees: []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}},
			}

			expected := map[string]any{
				"Employees": []any{
					map[string]any{"Name": "Alice", "Age": 30},
					map[string]any{"Name": "Bob", "Age": 25},
				},
			}

			assert.Equal(t, expected, ToMap(company))
		})
	})

	t.Run("slices", func(t *testing.T) {
		t.Run("converts slice to array", func(t *testing.T) {
			input := []int{1, 2, 3}
			expected := []any{1, 2, 3}
			assert.Equal(t, expected, ToMap(input))
		})

		t.Run("handles nested slices", func(t *testing.T) {
			input := [][]int{{1, 2}, {3, 4}}
			expected := []any{
				[]any{1, 2},
				[]any{3, 4},
			}
			assert.Equal(t, expected, ToMap(input))
		})
	})

	t.Run("maps", func(t *testing.T) {
		t.Run("converts map to string keys", func(t *testing.T) {
			input := map[string]int{"one": 1, "two": 2}
			expected := map[string]any{
				"one": 1,
				"two": 2,
			}
			assert.Equal(t, expected, ToMap(input))
		})

		t.Run("handles non-string keys", func(t *testing.T) {
			input := map[int]string{1: "one", 2: "two"}
			expected := map[string]any{
				"1": "one",
				"2": "two",
			}
			assert.Equal(t, expected, ToMap(input))
		})
	})
}

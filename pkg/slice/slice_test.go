package slice

import (
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		if !Contains(s, 3) {
			t.Error("expected 3 to be in slice")
		}

		if Contains(s, 6) {
			t.Error("expected 6 not to be in slice")
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := Reverse(s)
		if len(r) != len(s) {
			t.Errorf("expected length to be %d, got %d", len(s), len(r))
		}

		if r[0] != 5 {
			t.Error("expected first element to be 5")
		}

		if r[len(r)-1] != 1 {
			t.Error("expected last element to be 1")
		}
	})

	t.Run("Sort", func(t *testing.T) {
		s := []int{5, 4, 3, 2, 1}
		r := Sort(s)
		if len(r) != len(s) {
			t.Errorf("expected length to be %d, got %d", len(s), len(r))
		}

		if r[0] != 1 {
			t.Error("expected first element to be 1")
		}

		if r[len(r)-1] != 5 {
			t.Error("expected last element to be 5")
		}
	})

	t.Run("Map", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := Map(s, func(v int) int {
				return v * 2
			})
			if len(r) != len(s) {
				t.Errorf("expected length to be %d, got %d", len(s), len(r))
			}

			if r[0] != 2 {
				t.Error("expected first element to be 2")
			}

			if r[len(r)-1] != 10 {
				t.Error("expected last element to be 10")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := Map(s, func(v string) string {
				return v + v
			})
			if len(r) != len(s) {
				t.Errorf("expected length to be %d, got %d", len(s), len(r))
			}

			if r[0] != "aa" {
				t.Error("expected first element to be aa")
			}

			if r[len(r)-1] != "ee" {
				t.Error("expected last element to be ee")
			}

			rUpper := Map(s, strings.ToUpper)

			if rUpper[0] != "A" {
				t.Error("expected first element to be A")
			}
			if rUpper[len(rUpper)-1] != "E" {
				t.Error("expected last element to be E")
			}
		})
	})

	t.Run("Filter", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := Filter(s, func(v int) bool {
				return v%2 == 0
			})
			if len(r) != 2 {
				t.Errorf("expected length to be 2, got %d", len(r))
			}

			if r[0] != 2 {
				t.Error("expected first element to be 2")
			}

			if r[len(r)-1] != 4 {
				t.Error("expected last element to be 4")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := Filter(s, func(v string) bool {
				return v == "a" || v == "e"
			})
			if len(r) != 2 {
				t.Errorf("expected length to be 2, got %d", len(r))
			}

			if r[0] != "a" {
				t.Error("expected first element to be a")
			}

			if r[len(r)-1] != "e" {
				t.Error("expected last element to be e")
			}
		})
	})

	t.Run("Reduce", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := Reduce(s, 0, func(acc, v int) int {
				return acc + v
			})
			if r != 15 {
				t.Errorf("expected result to be 15, got %d", r)
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := Reduce(s, "", func(acc, v string) string {
				return acc + v
			})
			if r != "abcde" {
				t.Errorf("expected result to be abcde, got %s", r)
			}
		})
	})

	t.Run("Merge", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := Merge(s, s)
		if len(r) != len(s)*2 {
			t.Errorf("expected length to be %d, got %d", len(s)*2, len(r))
		}

		if r[0] != 1 {
			t.Error("expected first element to be 1")
		}
		if r[len(r)-1] != 5 {
			t.Error("expected last element to be 5")
		}
	})

	t.Run("All", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := All(s, func(v int) bool {
				return v > 0
			})
			if !r {
				t.Error("expected result to be true")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := All(s, func(v string) bool {
				return v != ""
			})
			if !r {
				t.Error("expected result to be true")
			}
		})
	})

	t.Run("Any", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := Any(s, func(v int) bool {
				return v == 3
			})
			if !r {
				t.Error("expected result to be true")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := Any(s, func(v string) bool {
				return v == "c"
			})
			if !r {
				t.Error("expected result to be true")
			}
		})
	})

	t.Run("None", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := None(s, func(v int) bool {
				return v == 0
			})
			if !r {
				t.Error("expected result to be true")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := None(s, func(v string) bool {
				return v == ""
			})
			if !r {
				t.Error("expected result to be true")
			}
		})
	})

	t.Run("Unique", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
			r := Unique(s)
			if len(r) != len(s)/2 {
				t.Errorf("expected length to be %d, got %d", len(s)/2, len(r))
			}

			if r[0] != 1 {
				t.Error("expected first element to be 1")
			}
			if r[len(r)-1] != 5 {
				t.Error("expected last element to be 5")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"}
			r := Unique(s)
			if len(r) != len(s)/2 {
				t.Errorf("expected length to be %d, got %d", len(s)/2, len(r))
			}

			if r[0] != "a" {
				t.Error("expected first element to be a")
			}
			if r[len(r)-1] != "e" {
				t.Error("expected last element to be e")
			}
		})
	})

	t.Run("UniqueBy", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
			r := UniqueBy(s, func(v int) int {
				return v
			})
			if len(r) != len(s)/2 {
				t.Errorf("expected length to be %d, got %d", len(s)/2, len(r))
			}

			if r[0] != 1 {
				t.Error("expected first element to be 1")
			}
			if r[len(r)-1] != 5 {
				t.Error("expected last element to be 5")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"}
			r := UniqueBy(s, func(v string) string {
				return v
			})
			if len(r) != len(s)/2 {
				t.Errorf("expected length to be %d, got %d", len(s)/2, len(r))
			}

			if r[0] != "a" {
				t.Error("expected first element to be a")
			}
			if r[len(r)-1] != "e" {
				t.Error("expected last element to be e")
			}
		})
	})

	t.Run("None", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r := None(s, func(v int) bool {
				return v == 0
			})
			if !r {
				t.Error("expected result to be true")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r := None(s, func(v string) bool {
				return v == ""
			})
			if !r {
				t.Error("expected result to be true")
			}
		})
	})

	t.Run("Find", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r, ok := Find(s, func(v int) bool {
				return v == 3
			})
			if !ok {
				t.Error("expected ok to be true")
			}
			if r != 3 {
				t.Error("expected result to be 3")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r, ok := Find(s, func(v string) bool {
				return v == "c"
			})
			if !ok {
				t.Error("expected ok to be true")
			}
			if r != "c" {
				t.Error("expected result to be c")
			}
		})
	})

	t.Run("FindIndex", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			r, ok := FindIndex(s, func(v int) bool {
				return v == 3
			})
			if !ok {
				t.Error("expected ok to be true")
			}
			if r != 2 {
				t.Error("expected result to be 2")
			}
		})

		t.Run("string", func(t *testing.T) {
			s := []string{"a", "b", "c", "d", "e"}
			r, ok := FindIndex(s, func(v string) bool {
				return v == "c"
			})
			if !ok {
				t.Error("expected ok to be true")
			}
			if r != 2 {
				t.Error("expected result to be 2")
			}
		})
	})
}

func BenchmarkSlice(b *testing.B) {
	b.Run("Contains", func(b *testing.B) {
		s := []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			Contains(s, 3)
		}
	})

	b.Run("Reverse", func(b *testing.B) {
		s := []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			Reverse(s)
		}
	})

	b.Run("Sort", func(b *testing.B) {
		s := []int{5, 4, 3, 2, 1}
		for i := 0; i < b.N; i++ {
			Sort(s)
		}
	})

	b.Run("Map", func(b *testing.B) {
		b.Run("int", func(b *testing.B) {
			s := []int{1, 2, 3, 4, 5}
			for i := 0; i < b.N; i++ {
				Map(s, func(v int) int {
					return v * 2
				})
			}
		})
	})

	b.Run("Merge", func(b *testing.B) {
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{6, 7, 8, 9, 10}
		for i := 0; i < b.N; i++ {
			Merge(s1, s2)
		}
	})

	b.Run("Filter", func(b *testing.B) {
		s := []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			Filter(s, func(v int) bool {
				return v%2 == 0
			})
		}
	})

	b.Run("Any", func(b *testing.B) {
		s := []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			Any(s, func(v int) bool {
				return v%2 == 0
			})
		}
	})
}

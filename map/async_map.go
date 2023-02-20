package m

import "sync"

/*
AsyncMap is a generic threadsafe map. It can be used as a drop-in replacement for sync.Map

Example

	AsyncMap := NewAsyncMap[string, int]()
	AsyncMap.Store("a", 1)
	AsyncMap.Store("b", 2)
	AsyncMap.Load("a") // 1, true
*/
type AsyncMap[T comparable, V any] struct {
	item sync.Map
}

// NewAsyncMap returns a new pointer to an asyncMap.
func NewAsyncMap[T comparable, V any]() *AsyncMap[T, V] {
	return &AsyncMap[T, V]{}
}

/*
Store sets the value for a key.

Example

	asyncMap := NewAsyncMap[string, int]()
	asyncMap.Store("a", 1)
*/
func (m *AsyncMap[T, V]) Store(key T, value V) {
	m.item.Store(key, value)
}

/*
Load returns the value and true if the key exists in the map, otherwise it returns the zero value and false.

Example

	asyncMap := NewAsyncMap[string, int]()
	asyncMap.Store("a", 1)
	asyncMap.Load("b") // 0, false
	asyncMap.Load("a") // 1, true
*/
func (m *AsyncMap[T, V]) Load(key T) (value V, ok bool) {
	v, ok := m.item.Load(key)
	if !ok {
		return value, ok
	}

	return v.(V), ok
}

/*
LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.

Example

	asyncMap := NewAsyncMap[string, int]()
	asyncMap.LoadOrStore("a", 1) // 1, false
	asyncMap.LoadOrStore("a", 2) // 1, true
*/
func (m *AsyncMap[T, V]) LoadOrStore(key T, value V) (V, bool) {
	v, loaded := m.item.LoadOrStore(key, value)

	return v.(V), loaded
}

/*
LoadAndDelete deletes the value for a key, returning the previous value if any.

Example

	asyncMap := NewAsyncMap[string, int]()
	asyncMap.Store("a", 1)
	asyncMap.LoadAndDelete("a") // 1, true
	asyncMap.LoadAndDelete("a") // 0, false
*/
func (m *AsyncMap[T, V]) LoadAndDelete(key T) (value V, ok bool) {
	v, ok := m.item.LoadAndDelete(key)
	if !ok {
		return value, ok
	}

	return v.(V), ok
}

/*
Delete deletes the value for a key. If the key does not exist, it does nothing.

Example

	asyncMap := NewAsyncMap[string, int]()
	asyncMap.Store("a", 1)
	asyncMap.Delete("a")
*/
func (m *AsyncMap[T, V]) Delete(key T) {
	m.item.Delete(key)
}

/*
Range calls fn sequentially for each key and value present in the map. If fn returns false, range stops the iteration.

Example

	asyncMap := NewAsyncMap[string, int]()
	asyncMap.Store("a", 1)
	asyncMap.Store("b", 2)
	fn := func(key string, value int) bool {
		fmt.Printf("%s: %d", key, value)
		return true
	}
	asyncMap.Range(fn)
*/
func (m *AsyncMap[T, V]) Range(fn func(key T, value V) bool) {
	m.item.Range(func(key, value any) bool {
		return fn(key.(T), value.(V))
	})
}

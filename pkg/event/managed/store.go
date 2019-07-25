package managed

import (
	"sync"
)

//Store is an interface providing methods for storing and loading data
type Store interface {
	Push(data interface{})
	// Pop removes data from store and returns it to caller
	Pop() interface{}
	Dispose()
} 

// InMemoryStore in memory implementation of store
type InMemoryStore struct {
	m *sync.Map
}

// NewInMemoryStore returns new InMemoryStore instance
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		m: &sync.Map{},
	}
}

// Push pushes data to strore
func (m *InMemoryStore) Push(data interface{}) {
	m.m.Store(data, data)
}

// Pop pops data from store
func (m *InMemoryStore) Pop() interface{} {
	var key, val interface{}
	m.m.Range(func(k, v interface{}) bool {
		key = k
		val = v
		return false
	})
	if key != nil {
		m.m.Delete(key)
	}
	return val
}

// Dispose releases resources used by store
func (m *InMemoryStore) Dispose() {
	// m.m = nil
}

// IsEmpty checks if store is empty
func (m *InMemoryStore) IsEmpty() bool {
	counter := 0
	f := func(k, v interface{}) bool {
		counter++
		return true
	}
	m.m.Range(f)
	return counter == 0
}

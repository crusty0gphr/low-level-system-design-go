package memorypool

import (
	"sync"
)

// MemPool represents a simple memory pool.
// pool is represented as a buffered channel
type MemPool[T any] struct {
	pool  chan *Object[T]
	mutex sync.Mutex
}

// Object represents an object stored in the pool.
type Object[T any] struct {
	value T
}

// New creates a new memory pool with the specified capacity.
func New[T any](capacity int) *MemPool[T] {
	return &MemPool[T]{
		pool: make(chan *Object[T], capacity),
	}
}

// Get retrieves the value from the pool.
// returns the default value of the generic type
func (p *MemPool[T]) Get() T {
	var empty T // default value for type T
	select {
	case obj := <-p.pool:
		return obj.value
	default:
		return empty
	}
}

// Put returns an object back to the pool.
// uses mutexes of thread safety
func (p *MemPool[T]) Put(val T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	select {
	case p.pool <- &Object[T]{
		value: val,
	}:
	default: // MemPool is full, discard the object.
	}
}

package inmemcache

import (
	"sync"
	"time"
)

// CacheItem is a generic struct for storing a cached value and its expiration time
// - Value: data being cached
// - Expiration: timestamp indicating when the cached item will expire
type CacheItem[T any] struct {
	Value      T
	Expiration int64
}

// Cache a generic, thread-safe cache
// - storage: Map storing cached items
// - mu: Mutex for synchronizing access
type Cache[T any] struct {
	storage map[string]CacheItem[T]
	mu      sync.RWMutex
}

// New creates new Cache
// can store any item of any type T
func New[T any]() *Cache[T] {
	c := &Cache[T]{
		storage: make(map[string]CacheItem[T]),
	}

	// starts the cleaner goroutine
	go c.cleanup()
	return c
}

// cleanup removes expired items from the cache
func (c *Cache[T]) cleanup() {
	for {
		// sleep for 1 minute before each cleanup cycle
		time.Sleep(1 * time.Minute)
		now := time.Now().UnixNano()
		c.mu.Lock() // write lock to safely modify the cache

		for key, item := range c.storage {
			if item.Expiration-now < 0 {
				delete(c.storage, key)
			}
		}
		c.mu.Unlock()
	}
}

// Set stores a value of type T in the cache associated with the given key
func (c *Cache[T]) Set(key string, value T, ttl time.Duration) {
	c.mu.Lock() // lock the cache to prevent concurrent access.
	defer c.mu.Unlock()

	c.storage[key] = CacheItem[T]{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixMilli(),
	}
}

// Get retrieves the value associated with the key from the cache, if not expired
// Returns the value of type T and true if found and valid;
// otherwise, a zero value of type T and false
func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var empty T // default value for type T
	item, found := c.storage[key]
	expired := item.Expiration-time.Now().UnixMilli() < 0

	if !found || expired {
		return empty, false
	}
	return item.Value, true
}

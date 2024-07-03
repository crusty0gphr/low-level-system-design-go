package inmemcache

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeCache() *Cache[string] {
	cache := New[string]()
	cache.Set("key1", "value1", 100*time.Millisecond)
	cache.Set("key2", "value2", 10*time.Millisecond)
	cache.Set("key3", "value3", 5*time.Millisecond)
	cache.Set("key4", "value4", 50*time.Millisecond)
	cache.Set("key5", "value5", 200*time.Millisecond)
	cache.Set("key6", "value6", 500*time.Millisecond)
	return cache
}

func TestExpiration(t *testing.T) {
	cache := makeCache()
	time.Sleep(500 * time.Millisecond)

	// must expire
	value, exist := cache.Get("key1")
	assert.False(t, exist)
	assert.Empty(t, value)

	// must not expire
	value, exist = cache.Get("key3")
	assert.True(t, exist)
	assert.NotEmpty(t, value)
}

func TestCleanup(t *testing.T) {
	cache := makeCache()
	time.Sleep(2 * time.Minute)

	fmt.Println(cache.storage)
}

func TestConcurrentReads(t *testing.T) {
	var wg sync.WaitGroup
	cache := makeCache()

	reader := func() {
		defer wg.Done()

		i := rand.IntN(7-1) + 1
		key := fmt.Sprintf("key%d", i)

		_, _ = cache.Get(key)
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go reader()
	}
	wg.Wait()
}

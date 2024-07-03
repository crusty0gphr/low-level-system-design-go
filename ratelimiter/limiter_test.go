package ratelimiter

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	limiter := New(1, 5)

	for i := 0; i < 10; i++ {
		allowed := limiter.Allow()
		assert.False(t, allowed)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		allowed := limiter.Allow()
		assert.True(t, allowed)
	}
}

func TestConcurrentExec(t *testing.T) {
	var wg sync.WaitGroup
	limiter := New(1, 5)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		go func() {
			defer wg.Done()
			_ = limiter.Allow()
		}()
	}
	wg.Wait()
}

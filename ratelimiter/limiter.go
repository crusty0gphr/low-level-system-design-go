package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate   int // rate at which tokens are replenished (tokens per second)
	burst  int // max tokens that can be accumulated
	tokens int // number of available tokens
	mu     sync.Mutex
	last   time.Time // the last time tokens were updated
}

// New returns an instance of RateLimiter
// - rate: tokens per second
// - burst: maximum tokens allowed
func New(rate, burst int) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		burst:  burst,
		tokens: burst,
		last:   time.Now(),
	}
}

// Allow calculates how much time has passed since the last token update
// calculates how many tokens should be added based on time passed since the last update and the rate
func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.last).Seconds()

	// how many tokens added during the elapsed time period
	// allows to dynamically adjust the number of available tokens based on the time update
	r.tokens = int(elapsed * float64(r.rate))

	if r.tokens > r.burst {
		r.tokens = r.burst
	}

	r.last = now

	// ensure that tokens does not exceed the burst limit
	if r.tokens > 0 {
		r.tokens--
		return true
	}

	return false
}

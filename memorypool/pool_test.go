package memorypool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	pool := New[int](1024)

	for i := 0; i < 10; i++ {
		pool.Put(i)
	}

	val := pool.Get()
	assert.Equal(t, 0, val)
}

package threadpool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	pool := Run(5)
	for i := 0; i < 10; i++ {
		i := i
		pool.Submit(func() {
			fmt.Println("Processing task", i)
		})
	}
	pool.Stop()
}

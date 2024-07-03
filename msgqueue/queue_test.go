package msgqueue

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	queue := New[int64](10)

	for i := 0; i < 10; i++ {
		queue.RegisterHandler(i, func(m Message[int64]) error {
			fmt.Printf("Message handled: %d", m.Payload)
			return nil
		})
	}

	go queue.Start()
	for i := 0; i < 10; i++ {
		queue.Enqueue(Message[int64]{
			ID:      i,
			Payload: int64(i),
		})
	}
}

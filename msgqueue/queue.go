package msgqueue

import (
	"sync"
	"time"
)

type HandlerFn[T any] func(Message[T]) error

// Message represents a data structure for messages with an ID and Payload
type Message[T any] struct {
	ID      int //  message identifier
	Payload T
}

// MessageQueue manages a queue of messages and their handlers
type MessageQueue[T any] struct {
	queue    chan Message[T]                // message queue
	handlers map[int]func(Message[T]) error // map of message ID to handler function
	mu       sync.Mutex
}

// New creates a new MessageQueue
// size defines the buffer size of the queue
func New[T any](size int) *MessageQueue[T] {
	return &MessageQueue[T]{
		queue:    make(chan Message[T], size),
		handlers: make(map[int]func(Message[T]) error),
	}
}

// RegisterHandler registers a handler function for a specific message ID
func (mq *MessageQueue[T]) RegisterHandler(id int, handler HandlerFn[T]) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.handlers[id] = handler // collect handlers
}

// Enqueue adds a message to the message queue
func (mq *MessageQueue[T]) Enqueue(msg Message[T]) {
	mq.queue <- msg // add the message to the queue
}

// Start begins processing messages from the queue and invoking their handlers
func (mq *MessageQueue[T]) Start() {
	for msg := range mq.queue {
		mq.mu.Lock()
		handler, exists := mq.handlers[msg.ID]
		mq.mu.Unlock()

		if exists {
			if err := handler(msg); err != nil {
				// retry enqueueing the message after 1 second if there's an error
				go func(m Message[T]) {
					time.Sleep(1 * time.Second)
					mq.Enqueue(m)
				}(msg)
			}
		}
	}
}

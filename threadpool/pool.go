package threadpool

import "sync"

// Minimum size for the thread pool
const minPoolSize = 1

// Task represents a function that can be executed in the thread pool
type Task func()

// ThreadPool manages a pool of goroutines to execute tasks concurrently
type ThreadPool struct {
	tasks    chan Task      // Channel to receive tasks
	stopChan chan struct{}  // Channel to signal stopping workers
	wg       sync.WaitGroup // WaitGroup to wait for all tasks to complete
}

// Run creates and starts a new ThreadPool with the specified size
// If size is 0, it defaults to minPoolSize
func Run(size uint) *ThreadPool {
	if size == 0 {
		size = minPoolSize
	}

	pool := &ThreadPool{
		tasks:    make(chan Task),
		stopChan: make(chan struct{}),
	}

	// Start worker goroutines based on the pool size
	for i := uint(0); i < size; i++ {
		go pool.worker()
	}

	return pool
}

// worker is a goroutine that executes tasks from the tasks channel
func (pool *ThreadPool) worker() {
	// Infinite loop to process tasks until stop signal is received
	for {
		select {
		case task := <-pool.tasks:
			task() // Execute the task
		case <-pool.stopChan:
			return // Exit the loop when stop signal is received
		}
	}
}

// Submit adds a new task to the thread pool for execution
func (pool *ThreadPool) Submit(task Task) {
	pool.wg.Add(1) // Increment WaitGroup counter
	pool.tasks <- func() {
		defer pool.wg.Done() // Decrement WaitGroup counter when task completes
		task()               // Execute the submitted task
	}
}

// Stop stops all worker goroutines in the thread pool and waits for all tasks to complete
func (pool *ThreadPool) Stop() {
	close(pool.stopChan) // Signal all workers to stop
	pool.wg.Wait()       // Wait for all tasks to complete
}

# Thread Pool

**Question**: Implement a thread pool in Go to manage a fixed number of worker goroutines to process tasks concurrently.

**Follow-up**: How would you gracefully shut down the thread pool, ensuring all tasks are completed?
___
**Answer**: Create a pool of worker goroutines that pick tasks from a channel. Use a sync.WaitGroup to wait for all tasks to complete. Signal workers to stop by closing the channel.
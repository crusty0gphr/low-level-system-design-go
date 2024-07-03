# Message Queue

**Question**: Design a simple message queue in Go for asynchronous task processing.

**Follow-up**: How would you ensure message delivery and handle message retries in case of failures? Provide a sample implementation.
___
**Answer**: Implement a channel-based queue to store messages. Use worker goroutines to process messages. To ensure delivery and retries, maintain a retry queue and re-enqueue failed messages with a delay.

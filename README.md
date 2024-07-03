# Low-level system design
Some low-level system design interview questions focused on creating specific components in Go to achieve various functionalities

### Table of contents
1.	**In-Memory Cache with Expiration**<br>
•	Question: How would you design an in-memory cache in Go that supports TTL (Time-to-Live) for cached items?<br>
•	Follow-up: How would you handle concurrent access to the cache? Provide a sample implementation.
2.	**Rate Limiter**<br>
•	Question: Design a rate limiter in Go that limits the number of requests a user can make to an API within a given time frame.<br>
•	Follow-up: How would you handle different rate limits for different users? Provide a sample implementation.
3.	**Thread Pool**<br>
•	Question: Implement a thread pool in Go to manage a fixed number of worker goroutines to process tasks concurrently.<br>
•	Follow-up: How would you gracefully shut down the thread pool, ensuring all tasks are completed?
4.	**Memory Pool**<br>
•	Question: Design a memory pool in Go for efficient memory allocation and deallocation. Explain its use cases.<br>
•	Follow-up: How would you manage fragmentation within the pool? Provide a sample implementation.
5.	**Message Queue**<br>
•	Question: Design a simple message queue in Go for asynchronous task processing.<br>
•	Follow-up: How would you ensure message delivery and handle message retries in case of failures? Provide a sample implementation.
6.	**Distributed Lock**<br>
•	Question: Design a distributed locking mechanism in Go using a central service like Redis.<br>
•	Follow-up: How would you handle lock expiration and ensure the lock is released if a process crashes? Provide a sample implementation.
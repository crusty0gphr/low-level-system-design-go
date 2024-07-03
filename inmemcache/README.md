# In-Memory Cache with Expiration

**Question**: How would you design an in-memory cache in Go that supports TTL (Time-to-Live) for cached items?

**Follow-up**: How would you handle concurrent access to the cache? Provide a sample implementation.
___

**Answer**: Design a `map` to store cache items along with their expiration times. Use a goroutine to periodically clean up expired items. For concurrent access, use `sync.RWMutex` to allow multiple reads and exclusive writes.
# Rate Limiter

**Question**: Design a rate limiter in Go that limits the number of requests a user can make to an API within a given time frame.

**Follow-up**: How would you handle different rate limits for different users? Provide a sample implementation.
___
**Answer**: Implement a token bucket algorithm, where tokens are added at a fixed rate and each request consumes a token. To handle different users, maintain a separate token bucket for each user.
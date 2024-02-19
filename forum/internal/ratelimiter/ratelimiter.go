package ratelimiter

import (
	"sync"
	"time"
)

// user-specific rate limitter
type RateLimiter struct {
	mu          sync.Mutex
	lastRequest time.Time
	interval    time.Duration
}

// new rate limiter with specific limit
func NewRateLimiter(rateLimit float64) *RateLimiter {
	// calculate the minimum interval for each request
	interval := time.Second / time.Duration(rateLimit)
	return &RateLimiter{
		interval: interval,
	}
}

// check if user allowed based on rate limit
func (rl *RateLimiter) Allow() bool {
	// ensure exclusive access
	rl.mu.Lock()
	defer rl.mu.Unlock()
	// get current time
	currentTime := time.Now()
	// calculate difference between two consecutive requests
	timeSinceLastRequest := currentTime.Sub(rl.lastRequest)
	// allow when time since last request bigger than interval
	if timeSinceLastRequest >= rl.interval {
		// set time of last request
		rl.lastRequest = currentTime
		return true
	}

	return false
}

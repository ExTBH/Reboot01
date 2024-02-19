// ratelimiter_test.go
package ratelimiter_test

import (
	"sandbox/internal/ratelimiter"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	// Create a rate limiter with a rate of 2 requests per second
	limiter := ratelimiter.NewRateLimiter(2)
	// simulate consequetive user requests
	for i := 0; i < 10; i++ {
		// get limiter bool value
		allow := limiter.Allow()
		if allow {
			t.Log("Allowed")
		} else {
			t.Errorf("Request %d was not allowed", i+1)
		}
		// simulate api request
		time.Sleep(100 * time.Millisecond)
	}
}

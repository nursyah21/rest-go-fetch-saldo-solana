package helper

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu     sync.Mutex
	limits map[string][]time.Time
	maxReq int
	window time.Duration
}

func NewRateLimiter(maxRequests int, windowTime time.Duration) *RateLimiter {
	return &RateLimiter{
		limits: make(map[string][]time.Time),
		maxReq: maxRequests,
		window: windowTime,
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	logs := rl.limits[ip]
	var newLogs []time.Time
	for _, t := range logs {
		if t.After(cutoff) {
			newLogs = append(newLogs, t)
		}
	}
	rl.limits[ip] = newLogs

	if len(newLogs) >= rl.maxReq {
		return false
	}

	rl.limits[ip] = append(newLogs, now)
	return true
}

func WithRateLimit(rl *RateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		if !rl.Allow(ip) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "rate limit exceeded",
			})
			return
		}
		next(w, r)
	}
}

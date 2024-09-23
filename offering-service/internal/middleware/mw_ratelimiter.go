package mw

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func RateLimiterMiddleware(rps float64, burst int, cleanupInterval time.Duration) func(http.Handler) http.Handler {
	clientLimiters := make(map[string]*RateLimiter)
	var mu sync.Mutex

	go func() {
		for now := range time.Tick(cleanupInterval) {
			mu.Lock()
			for ip, limiter := range clientLimiters {
				if now.Sub(limiter.lastSeen) > cleanupInterval {
					delete(clientLimiters, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr

			mu.Lock()
			limiter, exists := clientLimiters[ip]
			if !exists {
				limiter = &RateLimiter{
					limiter:  rate.NewLimiter(rate.Limit(rps), burst),
					lastSeen: time.Now(),
				}
				clientLimiters[ip] = limiter
			}
			limiter.lastSeen = time.Now()
			mu.Unlock()

			if !limiter.limiter.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
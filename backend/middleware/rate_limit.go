package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	limiter  *rateLimiter
	lastSeen time.Time
}

type rateLimiter struct {
	tokens     float64
	maxTokens  float64
	refillRate float64
	lastRefill time.Time
	mu         sync.Mutex
}

func newRateLimiter(maxTokens, refillRate float64) *rateLimiter {
	return &rateLimiter{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()
	rl.lastRefill = now

	rl.tokens += elapsed * rl.refillRate
	if rl.tokens > rl.maxTokens {
		rl.tokens = rl.maxTokens
	}

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}
	return false
}

// RateLimit creates an IP-based token-bucket rate limiter middleware.
// maxTokens: maximum burst capacity (e.g. 30 requests)
// refillRate: number of tokens refilled per second (e.g. 10 per second)
func RateLimit() gin.HandlerFunc {
	var mu sync.Mutex
	clients := make(map[string]*client)

	// Background cleanup routine to prevent memory leaks from inactive IPs
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			mu.Lock()
			for ip, c := range clients {
				if time.Since(c.lastSeen) > 5*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		cli, exists := clients[ip]
		if !exists {
			// Allow burst of 30 requests, refilling at 10 requests per second
			cli = &client{
				limiter: newRateLimiter(30, 10),
			}
			clients[ip] = cli
		}
		cli.lastSeen = time.Now()
		mu.Unlock()

		if !cli.limiter.allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please slow down.",
			})
			return
		}

		c.Next()
	}
}

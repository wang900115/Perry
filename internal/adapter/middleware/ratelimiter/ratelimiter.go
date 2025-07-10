package ratelimiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	"github.com/wang900115/utils/convert"
)

type RateLimiter struct {
	response       responser.Response
	limiter        *redis_rate.Limiter // limiter
	limitPerSecond int                 // 限制每秒幾次
}

func NewRateLimiter(response responser.Response, redis redis.Client, limitPerSecond int) *RateLimiter {
	return &RateLimiter{response: response, limiter: redis_rate.NewLimiter(redis), limitPerSecond: limitPerSecond}
}

func (rl RateLimiter) Middleware(c *gin.Context) {
	clientIP := c.ClientIP()
	if clientIP != "" {
		limit := redis_rate.PerSecond(rl.limitPerSecond) // 產生一個速率限制器
		res, err := rl.limiter.Allow(c, clientIP, limit)
		if err != nil {
			rl.response.SereverFail503(c, err)
			c.Abort()
			return
		}

		h := c.Writer.Header()
		h.Set("X-RateLimit-Limit", convert.FromInt64ToString(int64(rl.limitPerSecond)))              // 每秒允許上限
		h.Set("X-RateLimit-Remaining", convert.FromInt64ToString(int64(res.Remaining)))              // 剩餘次數
		h.Set("X-RateLimit-Reset", convert.FromInt64ToString(time.Now().Add(res.ResetAfter).Unix())) // 下次可用時間
		h.Set("Retry-After", convert.FromInt64ToString(int64(res.ResetAfter/time.Second)))           // 幾秒後可用
		if res.Allowed == 0 {
			rl.response.ClientFail429(c, err)
			c.Abort()
			return
		}
	}

	c.Next()
}

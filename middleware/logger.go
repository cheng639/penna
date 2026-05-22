package middleware

import (
	"github.com/gin-gonic/gin"
	"penna/config"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		requestID, _ := c.Get("request_id")
		binding, _ := c.Get("binding")
		uid, _ := c.Get("uid")
		config.Logger().Info().
			Str("request_id", requestID.(string)).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Any("authorization", c.GetHeader("Authorization")).
			Any("body", binding).
			Any("uid", uid).
			Str("ip", c.ClientIP()).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Msg("request")
	}
}

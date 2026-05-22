package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"penna/config"
	"penna/model/response"
	"runtime"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1024*1024)
				n := runtime.Stack(buf, false)
				log.Printf("stack: %s", string(buf[:n]))

				requestID, _ := c.Get("request_id")
				binding, _ := c.Get("binding")
				uid, _ := c.Get("uid")
				config.Logger().Error().
					Str("request_id", requestID.(string)).
					Str("method", c.Request.Method).
					Str("path", c.Request.URL.Path).
					Str("query", c.Request.URL.RawQuery).
					Any("authorization", c.GetHeader("Authorization")).
					Any("body", binding).
					Any("uid", uid).
					Str("ip", c.ClientIP()).
					Int("status", c.Writer.Status()).
					Err(errors.New(fmt.Sprintf("%v", err))).
					Str("stack", string(buf[:n])).
					Msg("panic recovered")

				c.JSON(http.StatusInternalServerError, response.Response{
					Code: response.ERROR,
					Msg:  fmt.Sprintf("%v", err),
				})
			}
		}()

		c.Next()
	}
}

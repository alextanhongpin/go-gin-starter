package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// RequestID is a middleware that injects the requestId from the header to be logged. c.MustGet("RequestId") returns the requestId.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Request.Header.Get("X-Request-Id")

		if reqID == "" {
			reqID = xid.New().String()
		}

		// Expose it to use in the application
		c.Set("RequestId", reqID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", reqID)
		c.Next()
	}
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Toggle takes a boolean to enable/disable an endpoint
func Toggle(on bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !on {
			code := http.StatusNotImplemented
			c.JSON(code, gin.H{
				"code":    code,
				"message": "This endpoint is not available",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

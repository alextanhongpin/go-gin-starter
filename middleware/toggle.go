package middleware

import (
	"net/http"

	"github.com/alextanhongpin/go-gin-starter/model"

	"github.com/gin-gonic/gin"
)

// Toggle takes a boolean to enable/disable an endpoint
func Toggle(on bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !on {
			c.JSON(http.StatusNotImplemented, model.ErrorResponse{
				Code:    http.StatusNotImplemented,
				Message: "This endpoint is not available",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

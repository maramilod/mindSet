package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestCostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("requestStartTime", time.Now())

		c.Next()
	}
}

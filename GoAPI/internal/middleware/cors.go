package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	// config := cors.DefaultConfig()
	// config.AllowHeaders = []string{"Authorization"}
	// config.ExposeHeaders = []string{"Refresh"}
	// config.AllowAllOrigins = true
	return cors.Default()
}

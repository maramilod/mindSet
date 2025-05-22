package middleware

import (
	"bytes"
	"io"
	"mind-set/config"
	log "mind-set/internal/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// CustomLogger
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		cost := time.Since(c.GetTime("requestStartTime"))
		if config.Config.AppEnv != "production" {
			//const MaxBodyBytes = int64(65536)
			// c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
			payload, _ := io.ReadAll(c.Request.Body)
			path := c.Request.URL.Path
			log.Logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", c.Request.URL.RawQuery),
				zap.String("body", string(payload)),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.String("cost", cost.String()),
				zap.String("response", blw.body.String()),
			)
		}
	}
}

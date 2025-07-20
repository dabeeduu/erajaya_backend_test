package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		fields := logrus.Fields{
			"status_code": statusCode,
			"method":      c.Request.Method,
			"latency":     latency,
			"path":        c.Request.URL.Path,
			"ip":          c.ClientIP(),
		}

		switch {
		case statusCode >= 500:
			log.WithFields(fields).Error("Server error")
		case statusCode >= 400:
			log.WithFields(fields).Warn("Client error")
		default:
			log.WithFields(fields).Info("Successful request")
		}

	}
}

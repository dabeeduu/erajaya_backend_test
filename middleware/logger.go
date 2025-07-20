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

		log.WithFields(logrus.Fields{
			"status_code": statusCode,
			"method":      c.Request.Method,
			"latency":     latency,
			"path":        c.Request.URL,
			"ip":          c.ClientIP(),
		}).Info("Incoming request")

	}
}

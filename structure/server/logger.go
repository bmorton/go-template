package server

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		line := log.WithFields(logrus.Fields{
			"request_id": c.Writer.Header().Get("X-Request-Id"),
			"client_ip":  c.ClientIP(),
			"start":      start,
			"end":        end,
			"latency":    latency,
			"method":     c.Request.Method,
			"status":     c.Writer.Status(),
			"path":       c.Request.URL.Path,
		})

		if errors := c.Errors.String(); errors != "" {
			line.WithFields(logrus.Fields{
				"error": c.Errors.String(),
			}).Warn("Request complete")
		} else {
			line.Info("Request complete")
		}

	}
}

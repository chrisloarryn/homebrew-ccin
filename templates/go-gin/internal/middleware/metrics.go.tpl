{{- if .GCPProject}}
package middleware

import (
	"time"

	"{{.ProjectName}}/internal/metrics"

	"github.com/gin-gonic/gin"
)

// MetricsMiddleware records metrics for HTTP requests
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Record metrics
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()

		metrics.RecordHTTPRequest(method, path, statusCode, duration)
	}
}
{{- else}}
package middleware

// Placeholder for middleware when GCP is not enabled
{{- end}}

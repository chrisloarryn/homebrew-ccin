{{- if .GCPProject}}
package middleware

import (
	"time"

	"{{.ProjectName}}/internal/metrics"

	"github.com/gofiber/fiber/v2"
)

// MetricsMiddleware records metrics for HTTP requests
func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Record metrics
		duration := time.Since(start)
		statusCode := c.Response().StatusCode()
		method := c.Method()
		path := c.Route().Path

		metrics.RecordHTTPRequest(method, path, statusCode, duration)

		return err
	}
}
{{- else}}
package middleware

import "github.com/gofiber/fiber/v2"

// MetricsMiddleware is a no-op when GCP is not enabled
func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
{{- end}}

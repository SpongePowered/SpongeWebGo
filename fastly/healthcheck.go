package fastly

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
	"time"
)

const (
	userAgentHeader      = "User-Agent"
	healthCheckUserAgent = "fastly (healthcheck)"
)

func (c *Cache) LogHandler() macaron.Handler {
	fallback := macaron.Logger()
	return func(ctx *macaron.Context) {
		// Skip logging for fastly health checks
		userAgent := ctx.Req.Header.Get(userAgentHeader)
		if !strings.HasSuffix(userAgent, healthCheckUserAgent) {
			if _, err := ctx.Invoke(fallback); err != nil {
				panic(err)
			}
			return
		}

		c.healthCheck = true

		ctx.Next()

		// Log failed requests
		if status := ctx.Resp.Status(); status != http.StatusOK {
			c.Log.Println("Failed fastly health check", ctx.Req.RequestURI, "with", status, http.StatusText(status))
		}
	}
}

func (c *Cache) verifyHealthCheck() {
	for range time.Tick(1 * time.Hour) {
		// Warn if no health checks were received
		if c.healthCheck {
			c.healthCheck = false
		} else {
			c.Log.Println("No fastly health checks were processed in the last hour")
		}
	}
}

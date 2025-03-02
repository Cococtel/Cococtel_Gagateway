package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Definir métricas globales de API
var (
	APIRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total de solicitudes HTTP recibidas",
		},
		[]string{"method", "path"},
	)

	APILatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Duración de las solicitudes HTTP",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

var APIErrors = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_requests_errors_total",
		Help: "Total de errores HTTP por código de estado",
	},
	[]string{"method", "path", "status"},
)

// Middleware para capturar métricas HTTP
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		// Obtener la ruta correctamente
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		if path == "/favicon.ico" {
			return
		}

		APIRequests.WithLabelValues(c.Request.Method, path).Inc()
		APILatency.WithLabelValues(c.Request.Method, path).Observe(duration)

		// Registrar errores si la respuesta es 4xx o 5xx
		if c.Writer.Status() >= 400 {
			APIErrors.WithLabelValues(c.Request.Method, path, fmt.Sprintf("%d", c.Writer.Status())).Inc()
		}
	}
}

// Registrar métricas en Prometheus
func InitMiddlewareMetrics() {
	prometheus.MustRegister(APIRequests, APILatency)
}

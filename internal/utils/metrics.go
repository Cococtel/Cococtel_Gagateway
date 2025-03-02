package utils

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Definición de métricas
var (
	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "service_request_duration_seconds",
			Help:    "Request latency in seconds",
			Buckets: prometheus.DefBuckets, // Usa los buckets por defecto de Prometheus
		},
		[]string{"repository", "method"}, // Mantén solo estas dos etiquetas
	)

	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_request_total",
			Help: "Total number of requests received",
		},
		[]string{"repository", "method", "status_code"}, // Etiquetas correctas
	)

	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_request_errors_total",
			Help: "Total number of failed requests",
		},
		[]string{"repository", "method", "status_code"}, // Mantén solo estas tres etiquetas
	)

	inFlightRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_in_flight_requests",
			Help: "Number of in-flight requests",
		},
		[]string{"repository"}, // Solo por servicio
	)
)

func InitMetrics() {
	// Registrar métricas en Prometheus
	prometheus.MustRegister(requestLatency, requestCount, errorCount, inFlightRequests)
}

// MeasureRequest mide el tiempo de una solicitud y registra la métrica adecuada.
func MeasureRequest(repository, method string, start time.Time, statusCode int, err error) {
	duration := time.Since(start).Seconds()

	// Registrar latencia
	requestLatency.WithLabelValues(repository, method).Observe(duration)

	// Registrar tráfico
	requestCount.WithLabelValues(repository, method, fmt.Sprintf("%d", statusCode)).Inc()

	// Registrar errores solo si ocurrió un error
	if err != nil {
		errorCount.WithLabelValues(repository, method, fmt.Sprintf("%d", statusCode)).Inc()
	}

	// Manejo de saturación: incrementar y decrementar el número de solicitudes en proceso
	inFlightRequests.WithLabelValues(repository).Inc()
	defer inFlightRequests.WithLabelValues(repository).Dec()
}

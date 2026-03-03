package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// metrics holds the Prometheus counters and histograms registered for this service.
type metrics struct {
	httpRequestsTotal    *prometheus.CounterVec
	httpRequestDuration  *prometheus.HistogramVec
}

// newMetrics registers and returns the Prometheus metrics for the server.
// All metrics are namespaced under "kubeplatform" to avoid collisions in a
// shared Prometheus instance.
func newMetrics() *metrics {
	return &metrics{
		httpRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "kubeplatform",
				Name:      "http_requests_total",
				Help:      "Total number of HTTP requests by method, path, and status code.",
			},
			[]string{"method", "path", "status"},
		),
		httpRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "kubeplatform",
				Name:      "http_request_duration_seconds",
				Help:      "HTTP request latency distributions by method and path.",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "path"},
		),
	}
}

// metricsHandler returns the standard Prometheus exposition handler.
// This endpoint is intended to be scraped by Prometheus and must not be
// exposed on a public Ingress.
func metricsHandler() http.Handler {
	return promhttp.Handler()
}

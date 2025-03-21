// Code generated by metricsgen. DO NOT EDIT.

package proxy

import (
	"github.com/cometbft/cometbft/libs/metrics/discard"
	prometheus "github.com/cometbft/cometbft/libs/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		MethodTimingSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "method_timing_seconds",
			Help:      "Timing for each ABCI method.",

			Buckets: []float64{.0001, .0004, .002, .009, .02, .1, .65, 2, 6, 25},
		}, append(labels, "method", "type")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		MethodTimingSeconds: discard.NewHistogram(),
	}
}

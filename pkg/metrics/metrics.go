package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	SessionDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "session_time",
		Help: "Current temperature of the CPU.",
	})
	SessionCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "live_session_count",
			Help: "Number of live sessions",
		},
		[]string{"session"},
	)

	MissingChannelCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "missing_channel",
			Help: "No channel for a given dimension",
		},
		[]string{"missing_channel"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(SessionDuration)
	prometheus.MustRegister(SessionCounter)
	prometheus.MustRegister(MissingChannelCounter)
}
